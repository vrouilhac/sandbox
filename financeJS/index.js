function computeAmountWithInterest(inputRate, amount, numOfPeriod) {
  const rate = inputRate/100;
  const fv = amount * Math.pow((1 + rate), numOfPeriod);
  return fv;
}

function computeInterest(amount, rate, periods) {
  const v = computeAmountWithInterest(rate/12, amount, periods);
  const interest = amount - v;
  return interest;
}

function computeAmortizationSchedule(startAmount, periods, rate, monthlyLoan) {
  let compoundInterest = 0;
  const accumulatedAmortization = Array
      .from(Array(periods))
      .map((_, index) => index)
      .reduce((acc, index) => {
	const value = computeInterest(startAmount - acc, rate, 1)
	compoundInterest += +value;
	const capital_amorti = monthlyLoan + value;
	console.log({ 
	  value: value, 
	  capital_amorti: capital_amorti,
	  monthlyLoan,
	  acc
	});
	return acc + capital_amorti;
      }, 0);
  return {
    compoundInterest: +compoundInterest.toFixed(2),
    capitalStillDue: +(startAmount - accumulatedAmortization).toFixed(2),
    accumulatedAmortization: +accumulatedAmortization.toFixed(2),
    annualLoanPayment: +(monthlyLoan*12).toFixed(2)
  };
}

const mensuality = 1401.087456373;
const rate = 1.4;
const amount = 293_139;
const resellingYear = 10;
const periods = 12;
const capitalStillDue = computeAmortizationSchedule(amount, periods, rate, mensuality);
console.log({ capitalStillDue });
