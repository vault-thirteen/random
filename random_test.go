// random_test.go.

////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019..2020 by Vault Thirteen.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
////////////////////////////////////////////////////////////////////////////////
//
// Web Site Address:	https://github.com/vault-thirteen.
//
////////////////////////////////////////////////////////////////////////////////

package random

import (
	"fmt"
	"math"
	"testing"
)

func Test_Uint(t *testing.T) {

	var countStatistics map[uint]int
	var deviationThresholdMax float64
	var err error
	var iterationsCount int
	var iterationsAvgCountPerUniqueValue int
	var randomValue uint
	var randomValueMin uint
	var randomValueMax uint
	var uniqueRandomValuesCount uint

	deviationThresholdMax = 0.1 // 10 %.
	randomValueMin = 1
	randomValueMax = 5
	uniqueRandomValuesCount = randomValueMax - randomValueMin + 1
	iterationsCount = 100000
	iterationsAvgCountPerUniqueValue = iterationsCount / int(uniqueRandomValuesCount)

	// Prepare the Statistics Holder.
	countStatistics = make(map[uint]int)
	for i := randomValueMin; i <= randomValueMax; i++ {
		countStatistics[i] = 0
	}

	// Create random Numbers.
	for i := 1; i <= iterationsCount; i++ {
		randomValue, err = Uint(randomValueMin, randomValueMax)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		if (randomValue < randomValueMin) ||
			(randomValue > randomValueMax) {
			t.Error("Random Value is out of Range")
			t.FailNow()
		}
		countStatistics[randomValue]++
	}

	// Inspect the Statistics...

	// 1. Check the total Count.
	iterationsCountGot := 0
	for _, count := range countStatistics {
		iterationsCountGot += count
	}
	if iterationsCountGot != iterationsCount {
		t.Error("Iterations Count Error")
		t.FailNow()
	}

	// 2. Check marginal Counts.
	countMin := math.MaxInt64
	countMax := 0
	for _, count := range countStatistics {
		if count < countMin {
			countMin = count
		}
		if count > countMax {
			countMax = count
		}
	}
	if (countMax < iterationsAvgCountPerUniqueValue) ||
		(countMin > iterationsAvgCountPerUniqueValue) {
		t.Error("Values are not truly random")
		t.FailNow()
	}
	deviationUp := float64(countMax-iterationsAvgCountPerUniqueValue) /
		float64(iterationsAvgCountPerUniqueValue)
	deviationDown := float64(iterationsAvgCountPerUniqueValue-countMin) /
		float64(iterationsAvgCountPerUniqueValue)
	if (deviationUp > deviationThresholdMax) ||
		(deviationDown > deviationThresholdMax) {
		t.Error("Values are not truly random")
		t.FailNow()
	}

	fmt.Println(
		countStatistics,
		iterationsAvgCountPerUniqueValue,
		countMin,
		deviationUp,
		countMax,
		deviationDown,
	)
}
