package goyeppp

import (
	"math"
	"testing"
)

const (
	arraySize = vLength * 16

	// Polynomial Coefficients
	c0   Yep64f = 1.53270461724076346
	c1   Yep64f = 1.45339856462100293
	c2   Yep64f = 1.21078763026010761
	c3   Yep64f = 1.46952786401453397
	c4   Yep64f = 1.34249847863665017
	c5   Yep64f = 0.75093174077762164
	c6   Yep64f = 1.90239336671587562
	c7   Yep64f = 1.62162053962810579
	c8   Yep64f = 0.53312230473555923
	c9   Yep64f = 1.76588453111778762
	c10  Yep64f = 1.31215699612484679
	c11  Yep64f = 1.49636144227257237
	c12  Yep64f = 1.52170011054112963
	c13  Yep64f = 0.83637497322280110
	c14  Yep64f = 1.12764540941736043
	c15  Yep64f = 0.65513628703807597
	c16  Yep64f = 1.15879020877781906
	c17  Yep64f = 1.98262901973751791
	c18  Yep64f = 1.09134643523639479
	c19  Yep64f = 1.92898634047221235
	c20  Yep64f = 1.01233347751449659
	c21  Yep64f = 1.89462732589369078
	c22  Yep64f = 1.28216239080886344
	c23  Yep64f = 1.78448898277094016
	c24  Yep64f = 1.22382217182612910
	c25  Yep64f = 1.23434674193555734
	c26  Yep64f = 1.13914782832335501
	c27  Yep64f = 0.73506235075797319
	c28  Yep64f = 0.55461432517332724
	c29  Yep64f = 1.51704871121967963
	c30  Yep64f = 1.22430234239661516
	c31  Yep64f = 1.55001237689160722
	c32  Yep64f = 0.84197209952298114
	c33  Yep64f = 1.59396169927319749
	c34  Yep64f = 0.97067044414760438
	c35  Yep64f = 0.99001960195021281
	c36  Yep64f = 1.17887814292622884
	c37  Yep64f = 0.58955609453835851
	c38  Yep64f = 0.58145654861350322
	c39  Yep64f = 1.32447212043555583
	c40  Yep64f = 1.24673632882394241
	c41  Yep64f = 1.24571828921765111
	c42  Yep64f = 1.21901343493503215
	c43  Yep64f = 1.89453941213996638
	c44  Yep64f = 1.85561626872427416
	c45  Yep64f = 1.13302165522004133
	c46  Yep64f = 1.79145993815510725
	c47  Yep64f = 1.59227069037095317
	c48  Yep64f = 1.89104468672467114
	c49  Yep64f = 1.78733894997070918
	c50  Yep64f = 1.32648559107345081
	c51  Yep64f = 1.68531055586072865
	c52  Yep64f = 1.08980909640581993
	c53  Yep64f = 1.34308207822154847
	c54  Yep64f = 1.81689492849547059
	c55  Yep64f = 1.38582137073988747
	c56  Yep64f = 1.04974901183570510
	c57  Yep64f = 1.14348742300966456
	c58  Yep64f = 1.87597730040483323
	c59  Yep64f = 0.62131555899466420
	c60  Yep64f = 0.64710935668225787
	c61  Yep64f = 1.49846610600978751
	c62  Yep64f = 1.07834176789680957
	c63  Yep64f = 1.69130785175832059
	c64  Yep64f = 1.64547687732258793
	c65  Yep64f = 1.02441150427208083
	c66  Yep64f = 1.86129006037146541
	c67  Yep64f = 0.98309038830424073
	c68  Yep64f = 1.75444578237500969
	c69  Yep64f = 1.08698336765112349
	c70  Yep64f = 1.89455010772036759
	c71  Yep64f = 0.65812118412299539
	c72  Yep64f = 0.62102711487851459
	c73  Yep64f = 1.69991208083436747
	c74  Yep64f = 1.65467704495635767
	c75  Yep64f = 1.69599459626992174
	c76  Yep64f = 0.82365682103308750
	c77  Yep64f = 1.71353437063595036
	c78  Yep64f = 0.54992984722831769
	c79  Yep64f = 0.54717367088443119
	c80  Yep64f = 0.79915543248858154
	c81  Yep64f = 1.70160318364006257
	c82  Yep64f = 1.34441280175456970
	c83  Yep64f = 0.79789486341474966
	c84  Yep64f = 0.61517383020710754
	c85  Yep64f = 0.55177400048576055
	c86  Yep64f = 1.43229889543908696
	c87  Yep64f = 1.60658663666266949
	c88  Yep64f = 1.78861146369896090
	c89  Yep64f = 1.05843250742401821
	c90  Yep64f = 1.58481799048208832
	c91  Yep64f = 1.70954313374718085
	c92  Yep64f = 0.52590070195022226
	c93  Yep64f = 0.92705074709607885
	c94  Yep64f = 0.71442651832362455
	c95  Yep64f = 1.14752795948077643
	c96  Yep64f = 0.89860175106926404
	c97  Yep64f = 0.76771198245570573
	c98  Yep64f = 0.67059202034800746
	c99  Yep64f = 0.53785922275590729
	c100 Yep64f = 0.82098327929734880
)

var (
	x     = make([]Yep64f, arraySize)
	coefs = []Yep64f{c0, c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12, c13, c14, c15, c16, c17, c18, c19, c20, c21, c22, c23, c24, c25, c26, c27, c28, c29, c30, c31, c32, c33, c34, c35, c36, c37, c38, c39, c40, c41, c42, c43, c44, c45, c46, c47, c48, c49, c50, c51, c52, c53, c54, c55, c56, c57, c58, c59, c60, c61, c62, c63, c64, c65, c66, c67, c68, c69, c70, c71, c72, c73, c74, c75, c76, c77, c78, c79, c80, c81, c82, c83, c84, c85, c86, c87, c88, c89, c90, c91, c92, c93, c94, c95, c96, c97, c98, c99, c100}
)

func init() {
	// initialize random array

	rng := new(WELL1024a)

	gyInit()

	status := rng.InitWELL2014a()
	checkStatus(status)

	/* Populate the array x with random data */
	status = rng.GenerateUniformS64fS64fV64fAcc64(0.0, 100.0, x)

	checkStatus(status)

	gyRelease()

}

// naivePolynomial compute naive polynomial product
func naivePolynomial(xv []Yep64f) []Yep64f {
	yv := make([]Yep64f, len(xv))
	for i := 0; i < arraySize; i++ {
		x := xv[i]
		y := c0 + x*(c1+x*(c2+x*(c3+x*(c4+x*(c5+x*(c6+x*(c7+x*(c8+x*(c9+x*(c10+x*(c11+x*(c12+x*(c13+x*(c14+x*(c15+x*(c16+x*(c17+x*(c18+x*(c19+x*(c20+x*(c21+x*(c22+x*(c23+x*(c24+x*(c25+x*(c26+x*(c27+x*(c28+x*(c29+x*(c30+x*(c31+x*(c32+x*(c33+x*(c34+x*(c35+x*(c36+x*(c37+x*(c38+x*(c39+x*(c40+x*(c41+x*(c42+x*(c43+x*(c44+x*(c45+x*(c46+x*(c47+x*(c48+x*(c49+x*(c50+x*(c51+x*(c52+x*(c53+x*(c54+x*(c55+x*(c56+x*(c57+x*(c58+x*(c59+x*(c60+x*(c61+x*(c62+x*(c63+x*(c64+x*(c65+x*(c66+x*(c67+x*(c68+x*(c69+x*(c70+x*(c71+x*(c72+x*(c73+x*(c74+x*(c75+x*(c76+x*(c77+x*(c78+x*(c79+x*(c80+x*(c81+x*(c82+x*(c83+x*(c84+x*(c85+x*(c86+x*(c87+x*(c88+x*(c89+x*(c90+x*(c91+x*(c92+x*(c93+x*(c94+x*(c95+x*(c96+x*(c97+x*(c98+x*(c99+x*c100)))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))
		yv[i] = y
	}

	return yv
}

func yepppPolynomial(xv []Yep64f) []Yep64f {

	y, status := EvalPolynomial64(coefs, xv)

	checkStatus(status)

	return y

}

// maxPolynomialError return the max error between two polynomial results set
func maxPolynomialError(x, y []Yep64f) Yep64f {
	var maxError Yep64f
	for i := 0; i < len(x); i++ {
		if x[i] == 0.0 {
			continue
		}

		error := Yep64f(math.Abs(float64((x[i] - y[i]))) / math.Abs(float64(x[i])))
		if error > maxError {
			maxError = error

		}
		return maxError
	}
	return maxError
}

// compare the difference between naive and yeppp polynomial calc
func TestPolynomial(t *testing.T) {
	naive := naivePolynomial(x)

	gyInit()

	yepppRes := yepppPolynomial(x)

	gyRelease()

	maxError := maxPolynomialError(naive, yepppRes)

	t.Logf("Polynomial max error = %g", maxError)
}

func BenchmarkPolynomial(b *testing.B) {

	gyInit()

	for i := 0; i < b.N; i++ {
		_ = yepppPolynomial(x)
	}

	gyRelease()

}

func BenchmarkNaivePolynomial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = naivePolynomial(x)
	}
}

func BenchmarkNaiveEntropy(b *testing.B) {

	var entropy Yep64f
	for j := 0; j < b.N; j++ {
		entropy = 0.0
		for i := arraySize - 1; i != 0; i-- {
			p := x[i]
			entropy -= (p * Yep64f(math.Log(float64(p))))
		}
	}

	b.Logf("BenchmarkNaiveEntropy Entropy: %g\n", entropy)
}

func BenchmarkEntroy(b *testing.B) {

	blockSize := 1024
	blockLength := 0

	gyInit()

	var entropy Yep64f
	for j := 0; j < b.N; j++ {
		entropy = 0.0
		for i := 0; i < arraySize; i += blockSize {
			blockLength = arraySize - i
			if blockLength > blockSize {
				blockLength = blockSize
			}

			logP, _ := Log(x[i : i+blockLength])

			dotProduct, _ := DotProductV64fV64fS64f(x[i:i+blockLength], logP)
			entropy -= dotProduct

		}
	}
	gyRelease()
	b.Logf("BenchmarkEntropy Entropy: %g\n", entropy)
}
