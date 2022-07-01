import { funding } from "./funding";
import { startChain } from "./helpers/utils";
import { staking } from "./staking";
import { alice, bob, charlie } from "./helpers/accounts";
import { delegation } from "./delegation";
describe("chain", () => {
  // disable timeout
  jest.setTimeout(24 * 60 * 60 * 1000);

  beforeAll(async () => {
    await startChain();
    await alice.init();
    await bob.init();
    await charlie.init();
  });

  // funding
  describe("Funding", funding);
  // staking
  describe("Staking", staking);
  // // delegation
  describe("Delegation", delegation);
  afterAll(() => {
    // stop local chain
    global.chain.kill();
  });
});
