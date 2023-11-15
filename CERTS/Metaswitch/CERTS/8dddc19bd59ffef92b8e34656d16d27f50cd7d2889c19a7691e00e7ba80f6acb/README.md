# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Rainbow Communications SHAKEN Cert 1820

Tested At: 15 Nov 23 15:51 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 78 day(s)\
Subject: CN=Rainbow Communications SHAKEN Cert 1820, O=Rainbow Communications, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/861edae589b34de5d112e3448589fd6c3d346e94

[View certificate details](https://understandingwebpki.com/?cert=MIICaDCCAg2gAwIBAgIQaSan%2Bz6NzX4cD%2Bsbo1SZzDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDIwMTE2NDYxM1oXDTI0MDIwMTE2NDYxM1owYDELMAkGA1UEBhMCVVMxHzAdBgNVBAoMFlJhaW5ib3cgQ29tbXVuaWNhdGlvbnMxMDAuBgNVBAMMJ1JhaW5ib3cgQ29tbXVuaWNhdGlvbnMgU0hBS0VOIENlcnQgMTgyMDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABG7sZhPy9lPSI0AajruJr42Pxsg8mrCL%2BR2Gy28m1Obnb8%2BIT8G%2BpAY%2FdAnxxG4zRzeBsjfTWYWIEMBSz5%2BKyJajgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDE4MjAwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUz6e5r1HZC6PhkHOQBNiQP4JYgGowHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSQAwRgIhAIo2WP0xbZiyKEdoiuQAHZQD383%2BTHFeZuJ2NBcletjeAiEAhtkK%2FS4gD4xSfPyrJ7dosIkjf%2FISYd6w%2F2ZpB1lqiXk%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 15 Nov 23 16:51 UTC