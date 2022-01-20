// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgAddNewPool } from "./types/ledger/tx";
import { MsgSetLeastBond } from "./types/ledger/tx";
import { MsgSetChainBondingDuration } from "./types/ledger/tx";
import { MsgSetCommission } from "./types/ledger/tx";
import { MsgSetReceiver } from "./types/ledger/tx";
import { MsgSetUnbondFee } from "./types/ledger/tx";
import { MsgLiquidityUnbond } from "./types/ledger/tx";
import { MsgClearCurrentEraSnapShots } from "./types/ledger/tx";
import { MsgSetUnbondCommission } from "./types/ledger/tx";
import { MsgRemovePool } from "./types/ledger/tx";
import { MsgSetInitBond } from "./types/ledger/tx";
import { MsgSetPoolDetail } from "./types/ledger/tx";
import { MsgSetEraUnbondLimit } from "./types/ledger/tx";


const types = [
  ["/stafiprotocol.stafihub.ledger.MsgAddNewPool", MsgAddNewPool],
  ["/stafiprotocol.stafihub.ledger.MsgSetLeastBond", MsgSetLeastBond],
  ["/stafiprotocol.stafihub.ledger.MsgSetChainBondingDuration", MsgSetChainBondingDuration],
  ["/stafiprotocol.stafihub.ledger.MsgSetCommission", MsgSetCommission],
  ["/stafiprotocol.stafihub.ledger.MsgSetReceiver", MsgSetReceiver],
  ["/stafiprotocol.stafihub.ledger.MsgSetUnbondFee", MsgSetUnbondFee],
  ["/stafiprotocol.stafihub.ledger.MsgLiquidityUnbond", MsgLiquidityUnbond],
  ["/stafiprotocol.stafihub.ledger.MsgClearCurrentEraSnapShots", MsgClearCurrentEraSnapShots],
  ["/stafiprotocol.stafihub.ledger.MsgSetUnbondCommission", MsgSetUnbondCommission],
  ["/stafiprotocol.stafihub.ledger.MsgRemovePool", MsgRemovePool],
  ["/stafiprotocol.stafihub.ledger.MsgSetInitBond", MsgSetInitBond],
  ["/stafiprotocol.stafihub.ledger.MsgSetPoolDetail", MsgSetPoolDetail],
  ["/stafiprotocol.stafihub.ledger.MsgSetEraUnbondLimit", MsgSetEraUnbondLimit],
  
];
export const MissingWalletError = new Error("wallet is required");

const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;

  const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgAddNewPool: (data: MsgAddNewPool): EncodeObject => ({ typeUrl: "/stafiprotocol.stafihub.ledger.MsgAddNewPool", value: data }),
    msgSetLeastBond: (data: MsgSetLeastBond): EncodeObject => ({ typeUrl: "/stafiprotocol.stafihub.ledger.MsgSetLeastBond", value: data }),
    msgSetChainBondingDuration: (data: MsgSetChainBondingDuration): EncodeObject => ({ typeUrl: "/stafiprotocol.stafihub.ledger.MsgSetChainBondingDuration", value: data }),
    msgSetCommission: (data: MsgSetCommission): EncodeObject => ({ typeUrl: "/stafiprotocol.stafihub.ledger.MsgSetCommission", value: data }),
    msgSetReceiver: (data: MsgSetReceiver): EncodeObject => ({ typeUrl: "/stafiprotocol.stafihub.ledger.MsgSetReceiver", value: data }),
    msgSetUnbondFee: (data: MsgSetUnbondFee): EncodeObject => ({ typeUrl: "/stafiprotocol.stafihub.ledger.MsgSetUnbondFee", value: data }),
    msgLiquidityUnbond: (data: MsgLiquidityUnbond): EncodeObject => ({ typeUrl: "/stafiprotocol.stafihub.ledger.MsgLiquidityUnbond", value: data }),
    msgClearCurrentEraSnapShots: (data: MsgClearCurrentEraSnapShots): EncodeObject => ({ typeUrl: "/stafiprotocol.stafihub.ledger.MsgClearCurrentEraSnapShots", value: data }),
    msgSetUnbondCommission: (data: MsgSetUnbondCommission): EncodeObject => ({ typeUrl: "/stafiprotocol.stafihub.ledger.MsgSetUnbondCommission", value: data }),
    msgRemovePool: (data: MsgRemovePool): EncodeObject => ({ typeUrl: "/stafiprotocol.stafihub.ledger.MsgRemovePool", value: data }),
    msgSetInitBond: (data: MsgSetInitBond): EncodeObject => ({ typeUrl: "/stafiprotocol.stafihub.ledger.MsgSetInitBond", value: data }),
    msgSetPoolDetail: (data: MsgSetPoolDetail): EncodeObject => ({ typeUrl: "/stafiprotocol.stafihub.ledger.MsgSetPoolDetail", value: data }),
    msgSetEraUnbondLimit: (data: MsgSetEraUnbondLimit): EncodeObject => ({ typeUrl: "/stafiprotocol.stafihub.ledger.MsgSetEraUnbondLimit", value: data }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
