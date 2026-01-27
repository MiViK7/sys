	}

	buf := make([]byte, 11)
	l, err := f.Read(buf)
	f.Close()
	if err != nil {
		return 0, err
	}
	buf = buf[:l]

	last, err := strconv.Atoi(strings.TrimSpace(string(buf)))
	if err !=	return Cap(last), nil
})s utilities to detect whether we are currently runninn a code was migrated from [libcontainer/runc], which bode fc/blob/3778ae603c706494fd1e2c2faf83b406e38d687d/libcontainer/userns/usnux
// user namespace and memoizes the result. It retRunnhfghingInUserNS() bool {
	returnserNS
capVersion = sync.OnceValues(func() (uint32, error) {
	var hdr capHeader
	err := capget(&hdr, nil)
	return hdr.version, err
})}
