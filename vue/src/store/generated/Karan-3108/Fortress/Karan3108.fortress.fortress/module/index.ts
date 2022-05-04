// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgApproveFortress } from "./types/fortress/tx";
import { MsgCancelFortress } from "./types/fortress/tx";
import { MsgRepayFortress } from "./types/fortress/tx";
import { MsgLiquidateFortress } from "./types/fortress/tx";
import { MsgRequestFortress } from "./types/fortress/tx";


const types = [
  ["/Karan3108.fortress.fortress.MsgApproveFortress", MsgApproveFortress],
  ["/Karan3108.fortress.fortress.MsgCancelFortress", MsgCancelFortress],
  ["/Karan3108.fortress.fortress.MsgRepayFortress", MsgRepayFortress],
  ["/Karan3108.fortress.fortress.MsgLiquidateFortress", MsgLiquidateFortress],
  ["/Karan3108.fortress.fortress.MsgRequestFortress", MsgRequestFortress],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

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
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgApproveFortress: (data: MsgApproveFortress): EncodeObject => ({ typeUrl: "/Karan3108.fortress.fortress.MsgApproveFortress", value: MsgApproveFortress.fromPartial( data ) }),
    msgCancelFortress: (data: MsgCancelFortress): EncodeObject => ({ typeUrl: "/Karan3108.fortress.fortress.MsgCancelFortress", value: MsgCancelFortress.fromPartial( data ) }),
    msgRepayFortress: (data: MsgRepayFortress): EncodeObject => ({ typeUrl: "/Karan3108.fortress.fortress.MsgRepayFortress", value: MsgRepayFortress.fromPartial( data ) }),
    msgLiquidateFortress: (data: MsgLiquidateFortress): EncodeObject => ({ typeUrl: "/Karan3108.fortress.fortress.MsgLiquidateFortress", value: MsgLiquidateFortress.fromPartial( data ) }),
    msgRequestFortress: (data: MsgRequestFortress): EncodeObject => ({ typeUrl: "/Karan3108.fortress.fortress.MsgRequestFortress", value: MsgRequestFortress.fromPartial( data ) }),
    
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
