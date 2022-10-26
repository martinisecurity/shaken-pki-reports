# STIR/SHAKEN CA Ecosystem Compliance
## Neustar

### Certificate aef2851bc3a7c530c6bcbf864edaeb635cb67155
Tested At: 2022-10-26 22:30:52 +0000 UTC\
Initial Validity Period: 3653 day(s)\
Remaining Validity Period: 3596 day(s)\
Subject: CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Inc a TransUnion company, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN Root CA, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US

Link: https://pki.tncp.textnow.com/prod/stir-shaken-textnow-cert.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIDSzCCAvGgAwIBAgIUDPrgzVZpiVxazrrwKND35tksogkwCgYIKoZIzj0EAwIwgYgxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEzMDEGA1UEAwwqTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBSb290IENBMB4XDTIyMDgzMDA2MzkwOVoXDTMyMDgzMDA2MzkwOVowgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluYyBhIFRyYW5zVW5pb24gY29tcGFueTEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0yMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEz75o5%2FrFenHCCu9TMbcCVT2gcsn04DgTQpgdcKnWENkpDjj4z48TjyMs1koK4VeIbdykiCOZ%2FoqwUJgmFChP3qOCATgwggE0MA8GA1UdEwEB%2FwQFMAMBAf8wHwYDVR0jBBgwFoAUFNWx4g9D2wES8o4AOlMKXL1z19AwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMA8GA1UdJQQIMAYGBFUdJQAwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMB0GA1UdDgQWBBRGj5%2BBtFSRVGG9nrq5bN6CtB20BTAOBgNVHQ8BAf8EBAMCAYYwCgYIKoZIzj0EAwIDSAAwRQIhAJCmnUZLbfvNhx0oZRzpuzTB7jmMK6N4HWVUPmWu8SadAiA4ynqmpGkeW1%2Bx307uTDkRXCOtg2zVwIM7%2BTx%2FS1pu6g%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| w_cp1_3_ca_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, and O can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_ca_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_sti_ca_extension_unknown | error | ATIS-1000080v4 | STI certificate shall not include extensions that are not specified |
| e_cp1_3_ca_key_usage_crl_sign | error | CPv1.3 | The model for managing and communicating the status of revoked certificates is in the form of a distributed Certificate Revocation List (CRL) that is maintained by the STI-PA as described in ATIS-1000080 |
| n_pki_ca_key_usage | notice | SHAKEN PKI Best Practice | For CA certificates, the Key Usage extension should contain a single key usage value of keyCertSign |

Generated: 26/10/2022 at 22:31:35