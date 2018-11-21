package app

import (
	"github.com/coschain/contentos-go/app/table"
	"github.com/coschain/contentos-go/common/constants"
	"github.com/coschain/contentos-go/prototype"
)

func mustSuccess(b bool, val string) {
	if !b {
		panic(val)
	}
}

type AccountCreateEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.AccountCreateOperation
}

type TransferEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.TransferOperation
}

type PostEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.PostOperation
}
type ReplyEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.ReplyOperation
}
type VoteEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.VoteOperation
}
type BpRegisterEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.BpRegisterOperation
}
type BpUnregisterEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.BpUnregisterOperation
}

type BpVoteEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.BpVoteOperation
}

type FollowEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.FollowOperation
}

type TransferToVestingEvaluator struct {
	BaseEvaluator
	ctx *ApplyContext
	op  *prototype.TransferToVestingOperation
}

func (ev *AccountCreateEvaluator) Apply() {
	op := ev.op
	creatorWrap := table.NewSoAccountWrap(ev.ctx.db, op.Creator)

	opAssert(creatorWrap.CheckExist(), "creator not exist ")

	opAssert(creatorWrap.GetBalance().Value >= op.Fee.Value, "Insufficient balance to create account.")

	// check auth accounts
	for _, a := range op.Owner.AccountAuths {
		tmpAccountWrap := table.NewSoAccountWrap(ev.ctx.db, a.Name)
		opAssert(tmpAccountWrap.CheckExist(), "owner auth account not exist")
	}
	for _, a := range op.Active.AccountAuths {
		tmpAccountWrap := table.NewSoAccountWrap(ev.ctx.db, a.Name)
		opAssert(tmpAccountWrap.CheckExist(), "active auth account not exist")
	}
	for _, a := range op.Posting.AccountAuths {
		tmpAccountWrap := table.NewSoAccountWrap(ev.ctx.db, a.Name)
		opAssert(tmpAccountWrap.CheckExist(), "posting auth account not exist")
	}

	// sub creator's fee
	originBalance := creatorWrap.GetBalance()
	opAssertE( originBalance.Sub( op.Fee ), "creator balance overflow" )
	opAssert( creatorWrap.MdBalance(originBalance), "")

	// create account
	newAccountWrap := table.NewSoAccountWrap(ev.ctx.db, op.NewAccountName)
	opAssertE(newAccountWrap.Create(func(tInfo *table.SoAccount) {
		tInfo.Name = op.NewAccountName
		tInfo.Creator = op.Creator
		tInfo.CreatedTime = ev.ctx.control.HeadBlockTime()
		tInfo.PubKey = op.MemoKey
		tInfo.Balance = prototype.NewCoin(0)
		tInfo.VestingShares = op.Fee.ToVest()
	}), "duplicate create account object")

	// create account authority
	authorityWrap := table.NewSoAccountAuthorityObjectWrap(ev.ctx.db, op.NewAccountName)
	opAssertE(authorityWrap.Create(func(tInfo *table.SoAccountAuthorityObject) {
		tInfo.Account = op.NewAccountName
		tInfo.Posting = op.Posting
		tInfo.Active = op.Active
		tInfo.Owner = op.Owner
		tInfo.LastOwnerUpdate = prototype.NewTimePointSec(0)
	}), "duplicate create account authority object")

	// sub dynamic glaobal properties's total fee
	ev.ctx.control.TransferToVest( op.Fee )
}

func (ev *TransferEvaluator) Apply() {
	op := ev.op

	// @ active_challenged
	fromWrap := table.NewSoAccountWrap(ev.ctx.db, op.From)
	toWrap := table.NewSoAccountWrap(ev.ctx.db, op.To)

	opAssert(toWrap.CheckExist(), "To account do not exist ")

	fBalance := fromWrap.GetBalance()
	tBalance := toWrap.GetBalance()

	opAssertE( fBalance.Sub( op.Amount ), "Insufficient balance to transfer.")
	opAssert( fromWrap.MdBalance(fBalance), "")

	opAssertE( tBalance.Add( op.Amount ), "balance overflow")
	opAssert( toWrap.MdBalance(tBalance), "")
}

func (ev *ReplyEvaluator) Apply() {
	op := ev.op
	cidWrap := table.NewSoPostWrap(ev.ctx.db, &op.Uuid)
	pidWrap := table.NewSoPostWrap(ev.ctx.db, &op.ParentUuid)

	opAssert(!cidWrap.CheckExist(), "post uuid exist")
	opAssert(pidWrap.CheckExist(), "parent uuid do not exist")

	opAssert(pidWrap.GetDepth()+1 < constants.POST_MAX_DEPTH, "reply depth error")

	opAssertE(cidWrap.Create(func(t *table.SoPost) {
		t.PostId = op.Uuid
		t.Tags = nil
		t.Title = ""
		t.Author = op.Owner
		t.Body = op.Content
		t.Created = ev.ctx.control.HeadBlockTime()
		t.LastPayout = prototype.NewTimePointSec(0) //TODO
		t.Depth = pidWrap.GetDepth() + 1
		t.Children = 0
		t.RootId = pidWrap.GetRootId()
		t.ParentId = constants.POST_INVALID_ID
		t.VoteCnt = 0

		// TODO  CreatedOrder / ReplyOrder sort
		// maybe should implement in plugin-services

	}), "create reply error")

	// Modify Parent Object
	opAssert(pidWrap.MdChildren(pidWrap.GetChildren()+1), "Modify Parent Children Error")
}

func (ev *PostEvaluator) Apply() {
	op := ev.op
	idWrap := table.NewSoPostWrap(ev.ctx.db, &op.Uuid)

	opAssert(!idWrap.CheckExist(), "post uuid exist")

	opAssertE(idWrap.Create(func(t *table.SoPost) {
		t.PostId = op.Uuid
		t.Tags = op.Tags
		t.Title = op.Title
		t.Author = op.Owner
		t.Body = op.Content
		t.Created = ev.ctx.control.HeadBlockTime()
		t.LastPayout = prototype.NewTimePointSec(0) //TODO
		t.Depth = 0
		t.Children = 0
		t.RootId = t.PostId
		t.ParentId = constants.POST_INVALID_ID
		t.VoteCnt = 0

		// TODO  CreatedOrder / ReplyOrder sort
		// maybe should implement in plugin-services

	}), "create post error")
}

func (ev *VoteEvaluator) Apply() {
	op := ev.op

	voterId := prototype.VoterId{Voter: op.Voter, PostId: op.Idx}
	vidWrap := table.NewSoVoteWrap(ev.ctx.db, &voterId)
	pidWrap := table.NewSoPostWrap(ev.ctx.db, &op.Idx)

	opAssert(!pidWrap.CheckExist(), "post invalid")
	opAssert(!vidWrap.CheckExist(), "vote info exist")

	opAssertE(vidWrap.Create(func(t *table.SoVote) {
		t.Voter = &voterId
		t.PostId = op.Idx
		t.VoteTime = ev.ctx.control.HeadBlockTime()
	}), "create voter object error")

	opAssert(pidWrap.MdVoteCnt(pidWrap.GetVoteCnt()+1), "set vote count error")
}

func (ev *BpRegisterEvaluator) Apply() {
	op := ev.op
	witnessWrap := table.NewSoWitnessWrap(ev.ctx.db, op.Owner)

	opAssert(witnessWrap.CheckExist(), "witness already exist")

	opAssertE(witnessWrap.Create(func(t *table.SoWitness) {
		t.Owner = op.Owner
		t.CreatedTime = ev.ctx.control.HeadBlockTime()
		t.Url = op.Url
		t.SigningKey = op.BlockSigningKey

		// TODO add others
	}), "add witness record error")
}

func (ev *BpUnregisterEvaluator) Apply() {
	// unregister op cost too much cpu time
	panic("not yet implement")

}

func (ev *BpVoteEvaluator) Apply() {
	op := ev.op

	voterAccount := table.NewSoAccountWrap(ev.ctx.db, op.Voter)
	voteCnt := voterAccount.GetBpVoteCount()

	voterId := &prototype.BpVoterId{Voter: op.Voter, Witness: op.Witness}
	witnessId := &prototype.BpWitnessId{Voter: op.Voter, Witness: op.Witness}
	vidWrap := table.NewSoWitnessVoteWrap(ev.ctx.db, voterId)

	witAccWrap := table.NewSoAccountWrap( ev.ctx.db, op.Voter )
	opAssert(witAccWrap.CheckExist(), "witness account do not exist ")

	witnessWrap := table.NewSoWitnessWrap(ev.ctx.db, op.Witness)

	if op.Cancel {
		opAssert(voteCnt > 0, "vote count must not be 0")
		opAssert(vidWrap.CheckExist(), "vote record not exist")
		opAssert(vidWrap.RemoveWitnessVote(), "remove vote record error")
		opAssert(witnessWrap.GetVoteCount() > 0, "witness data error")
		opAssert(witnessWrap.MdVoteCount(witnessWrap.GetVoteCount()-1), "set witness data error")
		opAssert(voterAccount.MdBpVoteCount(voteCnt-1), "set voter data error")
	} else {
		opAssert(voteCnt < constants.MAX_BP_VOTE_COUNT, "vote count exceeding")

		opAssertE(vidWrap.Create(func(t *table.SoWitnessVote) {
			t.VoteTime = ev.ctx.control.HeadBlockTime()
			t.VoterId = voterId
			t.WitnessId = witnessId
		}), "add vote record error")

		opAssert(voterAccount.MdBpVoteCount(voteCnt+1), "set voter data error")
		opAssert(witnessWrap.MdVoteCount(witnessWrap.GetVoteCount()+1), "set witness data error")
	}

}

func (ev *FollowEvaluator) Apply() {
	op := ev.op

	followeeWrap := table.NewSoAccountWrap(ev.ctx.db, op.Following)
	opAssert(followeeWrap.CheckExist(), "followee account do not exist ")

	// TODO
	// Follow relation update should be implement is plugin-services
}

func (ev *TransferToVestingEvaluator) Apply() {
	op := ev.op

	fidWrap := table.NewSoAccountWrap( ev.ctx.db, op.From)
	tidWrap := table.NewSoAccountWrap( ev.ctx.db, op.To)

	opAssert(tidWrap.CheckExist(), "to account do not exist")

	fBalance := fidWrap.GetBalance()
	tVests   := tidWrap.GetVestingShares()
	addVests := prototype.NewVest( op.Amount.Value )

	opAssertE( fBalance.Sub(op.Amount) , "balance not enough")
	opAssert ( fidWrap.MdBalance(fBalance), "set from new balance error")

	opAssertE( tVests.Add(addVests), "vests error" )
	opAssert ( tidWrap.MdVestingShares(tVests), "set to new vests error")

	ev.ctx.control.TransferToVest( op.Amount )
}
