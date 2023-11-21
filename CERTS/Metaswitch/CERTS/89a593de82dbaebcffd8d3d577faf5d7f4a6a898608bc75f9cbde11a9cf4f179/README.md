# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Foothills Rural Tel SHAKEN Cert 0406

Tested At: 21 Nov 23 01:22 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 252 day(s)\
Subject: CN=Foothills Rural Tel SHAKEN Cert 0406, O=Foothills Rural Tel, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/75cbed3f99e2082742446010d73f3f45e7ffbb67

[View certificate details](https://understandingwebpki.com/?cert=MIICYDCCAgegAwIBAgIQZFd0eu3VCNlRoUFcegy44jAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDczMDE3MDQyOFoXDTI0MDcyOTE3MDQyOFowWjELMAkGA1UEBhMCVVMxHDAaBgNVBAoME0Zvb3RoaWxscyBSdXJhbCBUZWwxLTArBgNVBAMMJEZvb3RoaWxscyBSdXJhbCBUZWwgU0hBS0VOIENlcnQgMDQwNjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABNmtNTzTDvB0niDvGCq7%2F7Mk%2FocZjQHeZfKM0oTxLcAs6In%2FkBoeP7Yg2eLzYJ%2B9IHdO6wJJhfAi86WCmP2c1b2jgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDA0MDYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUZIkAHMop%2FPysHvirfxO6Rgei1bkwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDRwAwRAIgLvWIlHbVB7o7rZnmVdKYjGpDmKQjXdsHmLbNX7pL%2BDMCIFmjq8c6lDCJJvW1F9p2Yq6MWcx2FAt%2FtAwP%2BPLPw6uj)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 01:55 UTC