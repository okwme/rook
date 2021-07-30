// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgMove } from "./types/rook/game/tx";
import { MsgBuild } from "./types/rook/game/tx";
import { MsgCreate } from "./types/rook/game/tx";
const types = [
    ["/rook.game.MsgMove", MsgMove],
    ["/rook.game.MsgBuild", MsgBuild],
    ["/rook.game.MsgCreate", MsgCreate],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgMove: (data) => ({ typeUrl: "/rook.game.MsgMove", value: data }),
        msgBuild: (data) => ({ typeUrl: "/rook.game.MsgBuild", value: data }),
        msgCreate: (data) => ({ typeUrl: "/rook.game.MsgCreate", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
