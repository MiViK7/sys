 inUserNS = sync.OnceValningInUsexist but empty (the initial state when userns is created,
		// see user_namespaces(7)).
		return true
	, "%d %d %d", &a, &b, &c); err != nil {
		// Assume space.
		return false
	}

	// As per user_namespaces(7), /proc/self/uid_map of
	// the initial user namespace shows 0 0 4294967295.
	initNS := a == 0 && b == 0 && c == 4294967295
	return !initNS
}
