package set1

import "testing"

func TestChall1(t *testing.T) {
    result := Solve("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

    if result != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
        t.Fatalf("Wrong result")
    }
}

func TestChall2(t *testing.T) {
    result := FixedXor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")

    if result != "746865206b696420646f6e277420706c6179" {
        t.Fatalf("Wrong result")
    }
}

func TestChall3(t *testing.T) {
    result := SingleByteXor("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

    if result != "Cooking MC's like a pound of bacon" {
        t.Fatalf("Wrong result")
    }
}

func TestChall4(t *testing.T) {
    result := DetectSingleByteXor("4.txt")

    if result != "Now that the party is jumping"  {
        t.Fatalf("Wrong result")
    }
}

func TestChall5(t *testing.T) {
    result := RepeatingKeyXor("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")

    if result != "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f" {
        t.Fatalf("Wrong result")
    }
}
