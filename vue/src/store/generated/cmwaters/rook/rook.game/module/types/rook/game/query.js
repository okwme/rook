/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal';
import * as Long from 'long';
import { State, Overview, Params } from '../../rook/game/game';
export const protobufPackage = 'rook.game';
const baseQueryGetGameStateRequest = { id: 0 };
export const QueryGetGameStateRequest = {
    encode(message, writer = Writer.create()) {
        if (message.id !== 0) {
            writer.uint32(8).uint64(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetGameStateRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetGameStateRequest };
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetGameStateRequest };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        return message;
    }
};
const baseQueryGetGameStateResponse = {};
export const QueryGetGameStateResponse = {
    encode(message, writer = Writer.create()) {
        if (message.gameState !== undefined) {
            State.encode(message.gameState, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetGameStateResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.gameState = State.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetGameStateResponse };
        if (object.gameState !== undefined && object.gameState !== null) {
            message.gameState = State.fromJSON(object.gameState);
        }
        else {
            message.gameState = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.gameState !== undefined && (obj.gameState = message.gameState ? State.toJSON(message.gameState) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetGameStateResponse };
        if (object.gameState !== undefined && object.gameState !== null) {
            message.gameState = State.fromPartial(object.gameState);
        }
        else {
            message.gameState = undefined;
        }
        return message;
    }
};
const baseQueryGetGameRequest = { id: 0 };
export const QueryGetGameRequest = {
    encode(message, writer = Writer.create()) {
        if (message.id !== 0) {
            writer.uint32(8).uint64(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetGameRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetGameRequest };
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetGameRequest };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        return message;
    }
};
const baseQueryGetGameResponse = { players: '' };
export const QueryGetGameResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.players) {
            writer.uint32(10).string(v);
        }
        if (message.overview !== undefined) {
            Overview.encode(message.overview, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetGameResponse };
        message.players = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.players.push(reader.string());
                    break;
                case 2:
                    message.overview = Overview.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetGameResponse };
        message.players = [];
        if (object.players !== undefined && object.players !== null) {
            for (const e of object.players) {
                message.players.push(String(e));
            }
        }
        if (object.overview !== undefined && object.overview !== null) {
            message.overview = Overview.fromJSON(object.overview);
        }
        else {
            message.overview = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.players) {
            obj.players = message.players.map((e) => e);
        }
        else {
            obj.players = [];
        }
        message.overview !== undefined && (obj.overview = message.overview ? Overview.toJSON(message.overview) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetGameResponse };
        message.players = [];
        if (object.players !== undefined && object.players !== null) {
            for (const e of object.players) {
                message.players.push(e);
            }
        }
        if (object.overview !== undefined && object.overview !== null) {
            message.overview = Overview.fromPartial(object.overview);
        }
        else {
            message.overview = undefined;
        }
        return message;
    }
};
const baseQueryGetParamsRequest = { version: 0 };
export const QueryGetParamsRequest = {
    encode(message, writer = Writer.create()) {
        if (message.version !== 0) {
            writer.uint32(8).uint32(message.version);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetParamsRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.version = reader.uint32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetParamsRequest };
        if (object.version !== undefined && object.version !== null) {
            message.version = Number(object.version);
        }
        else {
            message.version = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.version !== undefined && (obj.version = message.version);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetParamsRequest };
        if (object.version !== undefined && object.version !== null) {
            message.version = object.version;
        }
        else {
            message.version = 0;
        }
        return message;
    }
};
const baseQueryGetParamsResponse = {};
export const QueryGetParamsResponse = {
    encode(message, writer = Writer.create()) {
        if (message.params !== undefined) {
            Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetParamsResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.params = Params.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetParamsResponse };
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromJSON(object.params);
        }
        else {
            message.params = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetParamsResponse };
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromPartial(object.params);
        }
        else {
            message.params = undefined;
        }
        return message;
    }
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    GameState(request) {
        const data = QueryGetGameStateRequest.encode(request).finish();
        const promise = this.rpc.request('rook.game.Query', 'GameState', data);
        return promise.then((data) => QueryGetGameStateResponse.decode(new Reader(data)));
    }
    Game(request) {
        const data = QueryGetGameRequest.encode(request).finish();
        const promise = this.rpc.request('rook.game.Query', 'Game', data);
        return promise.then((data) => QueryGetGameResponse.decode(new Reader(data)));
    }
    Params(request) {
        const data = QueryGetParamsRequest.encode(request).finish();
        const promise = this.rpc.request('rook.game.Query', 'Params', data);
        return promise.then((data) => QueryGetParamsResponse.decode(new Reader(data)));
    }
}
var globalThis = (() => {
    if (typeof globalThis !== 'undefined')
        return globalThis;
    if (typeof self !== 'undefined')
        return self;
    if (typeof window !== 'undefined')
        return window;
    if (typeof global !== 'undefined')
        return global;
    throw 'Unable to locate global object';
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER');
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
