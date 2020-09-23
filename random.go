// random.go.

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

// Random Number Generator.

// This Packages offers a convenient Generator of random Integer Numbers.

package random

import (
	crand "crypto/rand"
	"errors"
	"math"
	"math/big"
)

// Errors.
const (
	ErrLimits   = "Limits Error"
	ErrOverflow = "Overflow"
)

// Creates a new random unsigned Integer Number in the [min;max] Interval.
func Uint(
	min uint,
	max uint,
) (result uint, err error) {

	// Fool Check.
	if min >= max {
		err = errors.New(ErrLimits)
		return
	}

	// Preparation.
	var offset uint = min
	var ránge uint = max - min

	// Unfortunately, the 'big' Library does not accept unsigned Integer Values.
	// Check the Limits.
	if ránge > math.MaxInt64-1 {
		err = errors.New(ErrOverflow)
		return
	}
	var crandMax *big.Int
	crandMax = big.NewInt(int64(ránge) + 1)

	// Create a uniform random Value in [0; crandMax).
	var crandRandomValue *big.Int
	crandRandomValue, err = crand.Int(crand.Reader, crandMax)
	if err != nil {
		return
	}

	if !crandRandomValue.IsUint64() {
		err = errors.New(ErrOverflow)
		return
	}
	result = uint(crandRandomValue.Uint64()) + offset
	return
}
