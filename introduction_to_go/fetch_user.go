func saveToRedis(u *User) error {
  return nil
}

func findFromRedis(id string) (*User, error) {
  return &User{}, nil
}
func findFromSql(id string) (*User, error) {
  return &User{}, nil
}

// START1 OMIT
func fetchUser(id string) (*User, error) {
  var u *User
  var err error

  done := make(chan *User)
  go func() {
    u, _ := findFromRedis(id)  // HL
    done <- u
  }()

  select {  // HL
  case u = <-done:  // HL
  case <-time.After(REDIS_TIMEOUT * time.Millisecond):  // HL
  }  // HL

  if u == nil {
    u, err = findFromSql(id)  // HL
    if err != nil {
      return nil, err
    }
    saveToRedis(u)  // HL
  }
  return u, nil
}
// END1 OMIT
// START2 OMIT
func GetUser(id string) (*User, error) {
  return userMap.Get(id, fetchUser)
}

func (m *UserMap) Get(id string, fetcher func(id string) (*User, error)) (*User, error) {
  if u, ok := m[id]; ok {
    return u, nil
  }

  u, err := fetcher(id)
  if err != nil {
    return nil, err
  }

  m[id] = u
  go func() {  // HL
    <-time.After(USER_CACHING_TIME * time.Millisecond)  // HL
    delete(m, id)  // HL
  }()  // HL

  return u, nil
}
// END2 OMIT