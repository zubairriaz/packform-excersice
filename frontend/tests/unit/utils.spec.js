import { formatTime, formatAmount } from "../../src/utils";

describe("Testing Utils Custom Functions", () => {
  it("Format Amount Format the currency value to Austalian Currency ", () => {
    const val = { value: 100 };
    const returnVal = formatAmount(val);

    expect(returnVal).toEqual("A$100.00");
  });

  it("Format Amount formats empty value to N/A", () => {
    const val = { value: "" };
    const returnVal = formatAmount(val);

    expect(returnVal).toEqual("N/A");
  });

  it("Format Time format the time correctly", () => {
    const returnVal = formatTime({value:new Date('2015-03-25, 14:10')})
    const expectedVal = 'Mar 25, 08:10 PM'

    expect(returnVal).toEqual(expectedVal);
  });
});
