  errc := make(chan error)

  go func() { // HL
    err := msg.save()
    errc <- err // HL
  }()

  go func() { // HL
    c, err := msg.fetchChannel()
    msg.setChannel(c)
    errc <- err // HL
  }()

  go func() { // HL
    u, err := msg.fetchUser()
    msg.setUser(u)
    errc <- err // HL
  }()

  err1, err2, err3 := <-errc, <-errc, <- errc // HL
  if err1 != nil || err2 != nil || err3 != nil { /* ... */ }