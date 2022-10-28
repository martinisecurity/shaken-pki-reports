# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate ccd437835c59c46742e20ffa132954e2127fbc42
Tested At: 2022-10-28 18:54:52 +0000 UTC\
Initial Validity Period: 2555 day(s)\
Remaining Validity Period: 2500 day(s)\
Subject: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Root CA US, OU=Sansay CA, O=Sansay Corporation, L=San Diego, ST=California, C=US

Link: http://sti.comsapi.com/258k/ca.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIC2zCCAoCgAwIBAgIUFLVfOAX18HsTtfiw3u0g8lFwPpwwCgYIKoZIzj0EAwIwgYoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlTYW4gRGllZ28xGzAZBgNVBAoMElNhbnNheSBDb3Jwb3JhdGlvbjESMBAGA1UECwwJU2Fuc2F5IENBMSEwHwYDVQQDDBhTSEFLRU4gU2Fuc2F5IFJvb3QgQ0EgVVMwHhcNMjIwOTAyMjA1MzA5WhcNMjkwODMxMjA1MzA5WjCBhTELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExGzAZBgNVBAoMElNhbnNheSBDb3Jwb3JhdGlvbjESMBAGA1UECwwJU2Fuc2F5IENBMTAwLgYDVQQDDCdTSEFLRU4gU2Fuc2F5IEludGVybWVkaWF0ZSBDQSBVUyBXRVNUIDEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARuhlHmOYoUiuAbIvxDHYUdDmCzFO4NLi4r47NjEzoDYDCdCKjWnCWHepTAG9PxCjNPT3GBwC2wsmLzFVyn8sj6o4HGMIHDMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwHwYDVR0jBBgwFoAUCq7%2FlvCbQaO9332%2FbdpFqOgEG7kwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMA8GA1UdEwEB%2FwQFMAMBAf8wDgYDVR0PAQH%2FBAQDAgKEMAoGCCqGSM49BAMCA0kAMEYCIQCO0QfDf%2Bz9Uu3vOm9ClE4BIdBvTIAVQ%2B%2FIKkulxneUtwIhAP3UpcrQ3RyvPVSvgrZxogOP00CFjZVmc4GUTY6xV4nn)


| Code | Type | Source | Details |
|------|------|--------|---------|
| w_pki_ca_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| n_pki_ca_key_usage | notice | SHAKEN PKI Best Practice | For CA certificates, the Key Usage extension should contain a single key usage value of keyCertSign |

Generated: 28/10/2022 at 18:55:01