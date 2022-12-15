# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Carolina West Wireless SHAKEN Cert 5932

Tested At: 15 Dec 22 18:33 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 543 day(s)\
Subject: CN=Carolina West Wireless SHAKEN Cert 5932, O=Carolina West Wireless, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/d9dd9bf73998fe328ba9447674492e8217d56dcf

[View certificate details](https://understandingwebpki.com/?cert=MIICZzCCAg2gAwIBAgIQSdecT7baK5UudVRwIoNLGzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDYxMTE2MTQyM1oXDTI0MDYxMDE2MTQyM1owYDELMAkGA1UEBhMCVVMxHzAdBgNVBAoMFkNhcm9saW5hIFdlc3QgV2lyZWxlc3MxMDAuBgNVBAMMJ0Nhcm9saW5hIFdlc3QgV2lyZWxlc3MgU0hBS0VOIENlcnQgNTkzMjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKXnc%2Fv0vMyETit9O8L5foIALi7PCQLP2wf62EGVy5DBVY0rP9m5oyH2D%2FbBDPWQOmrlxDQrnKjeGJXXcVFXvYajgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDU5MzIwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUg%2FAYzz%2FCbpX8tOgxBLTX72VAerMwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSAAwRQIgfyNcan63MjzlqUTwoWki8Wo9NsuO%2FZNSKWjeASga2EACIQDAp%2BrzAvo3mOVVfB2afaEqNpZ1lfPPxS%2B85Zp78F5pOQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 15 Dec 22 18:35 UTC