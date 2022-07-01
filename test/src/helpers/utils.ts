import { spawn, ChildProcessWithoutNullStreams } from "child_process";
import { NETWORK } from "./constants";
import KyveSDK from "@kyve/sdk";

export const sleep = (milliseconds: number) => {
  return new Promise((resolve) => setTimeout(resolve, milliseconds));
};
declare global {
  var chain: ChildProcessWithoutNullStreams;
}
export const startChain = (): Promise<ChildProcessWithoutNullStreams> => {
  return new Promise<ChildProcessWithoutNullStreams>((resolve, reject) => {
    console.log("Starting chain on localhost ...");
    console.log("This may take up to a minute");

    global.chain = spawn("ignite", ["chain", "serve", "--reset-once"]);

    chain.stdout.on("data", (data: Buffer) => {
      if (data.toString().includes(`Token faucet`)) {
        console.log(`Chain started`);
        setTimeout(() => resolve(chain), 1000);
      }
    });
    chain.stderr.on("data", () => {
      reject();
    });

    chain.on("close", () => {
      reject();
    });
  });
};
export const restartChain = () => {
  return new Promise<ChildProcessWithoutNullStreams>((resolve) => {
    global.chain.on("close", function test() {
      setTimeout(async () => {
        await startChain();
        global.chain.removeAllListeners("close");
        resolve(global.chain);
      }, 1000);
    });
    global.chain.kill();
  });
};

export const localSdk = new KyveSDK(NETWORK);
export const lcdClient = localSdk.createLCDClient();
