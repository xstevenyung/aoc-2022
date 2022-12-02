import fs from "fs/promises";
import path from "path";

async function main() {
  const data = await fs.readFile(path.join(__dirname, "./input"), "utf-8");

  let supplyByElves: Record<number, number> = {};
  let currentIndex = 0;
  for (const value of data.split("\n")) {
    if (value === "") {
      currentIndex += 1;
      continue;
    }

    supplyByElves[currentIndex] = (supplyByElves[currentIndex] || 0) + +value;
  }

  return Math.max(...Object.values(supplyByElves));
}

main().then(console.log);
