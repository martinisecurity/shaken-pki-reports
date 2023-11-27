# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Rainbow Communications SHAKEN Cert 1820

Tested At: 27 Nov 23 22:21 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 66 day(s)\
Subject: CN=Rainbow Communications SHAKEN Cert 1820, O=Rainbow Communications, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/861edae589b34de5d112e3448589fd6c3d346e94

[View certificate details](https://understandingwebpki.com/?cert=MIICaDCCAg2gAwIBAgIQaSan%2Bz6NzX4cD%2Bsbo1SZzDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDIwMTE2NDYxM1oXDTI0MDIwMTE2NDYxM1owYDELMAkGA1UEBhMCVVMxHzAdBgNVBAoMFlJhaW5ib3cgQ29tbXVuaWNhdGlvbnMxMDAuBgNVBAMMJ1JhaW5ib3cgQ29tbXVuaWNhdGlvbnMgU0hBS0VOIENlcnQgMTgyMDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABG7sZhPy9lPSI0AajruJr42Pxsg8mrCL%2BR2Gy28m1Obnb8%2BIT8G%2BpAY%2FdAnxxG4zRzeBsjfTWYWIEMBSz5%2BKyJajgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDE4MjAwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUz6e5r1HZC6PhkHOQBNiQP4JYgGowHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSQAwRgIhAIo2WP0xbZiyKEdoiuQAHZQD383%2BTHFeZuJ2NBcletjeAiEAhtkK%2FS4gD4xSfPyrJ7dosIkjf%2FISYd6w%2F2ZpB1lqiXk%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 27 Nov 23 22:56 UTC