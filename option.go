package hoi

type Options struct {
	Server  bool `short:"s" long:"server" description:"Start hoi server."`
	Clear   bool `short:"c" long:"clear" description:"Clear all symlinks by removing contents under public directory."`
	Version bool `long:"version" description:"Show version"`
}
