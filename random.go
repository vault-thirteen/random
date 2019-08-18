////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019 by Vault Thirteen.
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

// Random Number Generator.

// This Packages offers a convenient Generator of random Integer Numbers.

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

	var crandMax *big.Int
	var crandRandomValue *big.Int
	var offset uint
	var ránge uint

	// Fool Check.
	if min >= max {
		err = errors.New(ErrLimits)
		return
	}

	// Preparation.
	offset = min
	ránge = max - min

	// Unfortunately, the 'big' Library does not accept unsigned Integer Values.
	// Check the Limits.
	if ránge > math.MaxInt64-1 {
		err = errors.New(ErrOverflow)
		return
	}
	crandMax = big.NewInt(int64(ránge) + 1)

	// Create a uniform random Value in [0; crandMax).
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
