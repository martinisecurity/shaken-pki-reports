# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Foothills Rural Tel SHAKEN Cert 0406

Tested At: 12 Apr 23 21:55 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 474 day(s)\
Subject: CN=Foothills Rural Tel SHAKEN Cert 0406, O=Foothills Rural Tel, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/5fe148da9e593bacaa348da816e2104790886dbc

[View certificate details](https://understandingwebpki.com/?cert=MIICYDCCAgegAwIBAgIQZFd0eu3VCNlRoUFcegy44jAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDczMDE3MDQyOFoXDTI0MDcyOTE3MDQyOFowWjELMAkGA1UEBhMCVVMxHDAaBgNVBAoME0Zvb3RoaWxscyBSdXJhbCBUZWwxLTArBgNVBAMMJEZvb3RoaWxscyBSdXJhbCBUZWwgU0hBS0VOIENlcnQgMDQwNjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABNmtNTzTDvB0niDvGCq7%2F7Mk%2FocZjQHeZfKM0oTxLcAs6In%2FkBoeP7Yg2eLzYJ%2B9IHdO6wJJhfAi86WCmP2c1b2jgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDA0MDYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUZIkAHMop%2FPysHvirfxO6Rgei1bkwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDRwAwRAIgLvWIlHbVB7o7rZnmVdKYjGpDmKQjXdsHmLbNX7pL%2BDMCIFmjq8c6lDCJJvW1F9p2Yq6MWcx2FAt%2FtAwP%2BPLPw6uj)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 12 Apr 23 22:02 UTC