# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Neustar Certified Caller ID SHAKEN CA-2

Tested At: 31 Oct 22 16:43 UTC\
Initial Validity Period: 3653 day(s)\
Remaining Validity Period: 3591 day(s)\
Subject: CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Inc a TransUnion company, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN Root CA, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US

View: [Click to view](https://understandingwebpki.com/?cert=MIIDSzCCAvGgAwIBAgIUDPrgzVZpiVxazrrwKND35tksogkwCgYIKoZIzj0EAwIwgYgxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEzMDEGA1UEAwwqTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBSb290IENBMB4XDTIyMDgzMDA2MzkwOVoXDTMyMDgzMDA2MzkwOVowgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluYyBhIFRyYW5zVW5pb24gY29tcGFueTEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0yMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEz75o5%2FrFenHCCu9TMbcCVT2gcsn04DgTQpgdcKnWENkpDjj4z48TjyMs1koK4VeIbdykiCOZ%2FoqwUJgmFChP3qOCATgwggE0MA8GA1UdEwEB%2FwQFMAMBAf8wHwYDVR0jBBgwFoAUFNWx4g9D2wES8o4AOlMKXL1z19AwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMA8GA1UdJQQIMAYGBFUdJQAwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMB0GA1UdDgQWBBRGj5%2BBtFSRVGG9nrq5bN6CtB20BTAOBgNVHQ8BAf8EBAMCAYYwCgYIKoZIzj0EAwIDSAAwRQIhAJCmnUZLbfvNhx0oZRzpuzTB7jmMK6N4HWVUPmWu8SadAiA4ynqmpGkeW1%2Bx307uTDkRXCOtg2zVwIM7%2BTx%2FS1pu6g%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_cp1_3_ca_key_usage_crl_sign](../../ISSUES/e_cp1_3_ca_key_usage_crl_sign/README.md) | error | United States SHAKEN CP | The model for managing and communicating the status of revoked certificates is in the form of a distributed Certificate Revocation List (CRL) that is maintained by the STI-PA as described in ATIS-1000080 |
| [e_sti_ca_extension_unknown](../../ISSUES/e_sti_ca_extension_unknown/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall not include extensions that are not specified |
| [n_pki_ca_key_usage](../../ISSUES/n_pki_ca_key_usage/README.md) | notice | SHAKEN PKI Best Practice | For CA certificates, the Key Usage extension should contain a single key usage value of keyCertSign |
| [w_pki_ca_subject_rdn_unknown](../../ISSUES/w_pki_ca_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 31/10/2022 at 16:43:22