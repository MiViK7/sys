s) (ut32, error) {
	var hdr capHeader
ideáerr := capget(&hdr, nil)
	return hdr.version, er
