package main

import (
  "crypto/rand"
  "crypto/sha256"
  "fmt"
  "encoding/binary"
)

const BLOCK_SIZE int = 32
const HASH_SIZE int = 256

func main() {
  message := "Hi"
  fmt.Println(message)
  lamport := GenerateKeys()
  signedMessage := lamport.Sign([]byte(message))
  fmt.Println(signedMessage)
}

func generatePrivateKey() ([HASH_SIZE][]byte, [HASH_SIZE][]byte) {
    var set_1 [HASH_SIZE][]byte
    var set_2 [HASH_SIZE][]byte
    for counter := 0; counter < HASH_SIZE; counter++ {
      token_1 := make([]byte, BLOCK_SIZE)
      token_2 := make([]byte, BLOCK_SIZE)
      _, err := rand.Read(token_1)
      if err != nil {
        fmt.Println(err)
      }
      _, err = rand.Read(token_2)
      if err != nil {
        fmt.Println(err)
      }
      set_1[counter] = token_1
      set_2[counter] = token_2
    }
    return set_1, set_2
}

func generatePublicKey(secretKey_1 [HASH_SIZE][]byte,
                        secretKey_2 [HASH_SIZE][]byte) ([HASH_SIZE][]byte, [HASH_SIZE][]byte) {
  var publicKey_1, publicKey_2 [HASH_SIZE][]byte
  for counter := 0; counter < HASH_SIZE; counter++ {
    hasher := sha256.New()
    hasher.Write(secretKey_1[counter])
    publicKey_1[counter] = hasher.Sum(nil)
    hasher.Write(secretKey_2[counter])
    publicKey_2[counter] = hasher.Sum(nil)
  }
  return publicKey_1, publicKey_2
}

func GenerateKeys() Lamport {
  secretKey_1, secretKey_2 := generatePrivateKey()
  publicKey_1, publicKey_2 := generatePublicKey(secretKey_1, secretKey_2)
  return CreateLamportSignature(secretKey_1, secretKey_2, publicKey_1, publicKey_2)
}

// Lamport stuff - TODO: place in diff file//
type Lamport struct {
  secretKey_1 [HASH_SIZE][]byte
  secretKey_2 [HASH_SIZE][]byte
  PublicKey_1 [HASH_SIZE][]byte
  PublicKey_2 [HASH_SIZE][]byte
}

func CreateLamportSignature(sk1 [HASH_SIZE][]byte,
                            sk2 [HASH_SIZE][]byte,
                            pk1 [HASH_SIZE][]byte,
                            pk2 [HASH_SIZE][]byte) Lamport {
  lamportSig := Lamport{}
  lamportSig.secretKey_1 = sk1;
  lamportSig.secretKey_2 = sk2;
  lamportSig.PublicKey_1 = pk1;
  lamportSig.PublicKey_2 = pk2;
  return lamportSig
}

func (l *Lamport) Sign(message []byte) []byte {
  hasher := sha256.New()
  hasher.Write(message)
  hashedMessage := hasher.Sum(nil)
  var signedMessage []byte
  var byteVal byte
  for counter := 0; counter < len(hashedMessage); counter++ {
    currNum = hashedMessage[counter]
    // bit operations to determine which to use
    for bit := 0; bit < 8; bit++ {
        bitVal = currNum & (2**bit)
        if bitVal == 0 {
          // take from group one
        } else {
          // take from group two
        }
    }

    // go through each bit, determine if you need to take from array 1 or array 2
    // go to the associated
  }
  return hashedMessage
}

func getBinaryForm(deci int) {
  return binary.LittleEndian.Uint32(deci)
}
