import fs from "node:fs/promises";
import { yamlParse } from "yaml-cfn";

test("cfn.yml matches snapshot", async () => {
  const file = await fs.readFile(
    new URL("cfn.yml", import.meta.url),
    "utf8"
  );

  const doc = yamlParse(file);
  expect(doc).toMatchSnapshot();
});
