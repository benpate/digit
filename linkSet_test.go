package digit

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLinkSetCreate(t *testing.T) {
	set := NewLinkSet(10)
	require.Equal(t, 0, len(set))
}

func TestLinkSet(t *testing.T) {

	set := NewLinkSet(0)

	{
		link := NewLink("friend", "text/html", "http://example.com/friend")
		set.ApplyBy("rel", link)

		require.Equal(t, 1, len(set))
		require.Equal(t, link, set.FindBy("rel", "friend"))
		require.Equal(t, link, set.FindBy("type", "text/html"))
		require.Equal(t, "http://example.com/friend", set.FindBy("rel", "friend").GetString("href"))
	}
	{
		link := NewLink("parent", "application/json", "http://example.com/parent")
		set.ApplyBy("rel", link)

		require.Equal(t, 2, len(set))
		require.Equal(t, link, set.FindBy("rel", "parent"))
		require.Equal(t, link, set.FindBy("type", "application/json"))
	}
	{
		link := NewLink("sibling", "text/markdown", "http://example.com/sibling")
		set.ApplyBy("rel", link)

		require.Equal(t, 3, len(set))
		require.Equal(t, link, set.FindBy("rel", "sibling"))
		require.Equal(t, link, set.FindBy("type", "text/markdown"))
	}
	{
		link := NewLink("friend", "text/html", "http://example.com/friend-but-a-different-one")
		set.ApplyBy("rel", link)

		require.Equal(t, 3, len(set))
		require.Equal(t, link, set.FindBy("rel", "friend"))
		require.Equal(t, link, set.FindBy("type", "text/html"))
		require.Equal(t, "http://example.com/friend-but-a-different-one", set.FindBy("rel", "friend").GetString("href"))
	}

	{
		link := set.FindBy("rel", "nobody")
		require.True(t, link.IsEmpty())
	}

	set.RemoveBy("rel", "friend")
	require.Equal(t, 2, len(set))
}
