# STIR/SHAKEN CA Ecosystem Compliance

## Certificate GVTC SHAKEN Cert 2083

Tested At: 28 Nov 23 19:52 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 817 day(s)\
Subject: CN=GVTC SHAKEN Cert 2083, O=GVTC, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/90b8c01a8cc0c06e82b9517d599d24c14dbb859b

[View certificate details](https://understandingwebpki.com/?cert=MIICQzCCAemgAwIBAgIQWvc%2BUVymE9%2FSCN5C7oycxDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDIyMjIwNDU1MFoXDTI2MDIyMTIwNDU1MFowPDELMAkGA1UEBhMCVVMxDTALBgNVBAoMBEdWVEMxHjAcBgNVBAMMFUdWVEMgU0hBS0VOIENlcnQgMjA4MzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABADtLAkxWmg4BW5w%2FbMBLJD9WkM8CzDhKhUOgu3sLuzg1BGx0XeXK8ljM1KAqvzQkgInk5j0fYz38INxvl%2BcG9ajgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDIwODMwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUCX%2FAxX5zTjc8vzedB0ag14YoJOEwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDSAAwRQIgBCYJKpLTexw2OFKdS9O%2F3n%2B0s9EwwFdNt9FnO753X94CIQDcZF0YCSNsAJemwCFqvgEs3rcOHI%2FEnplFCUe8zXbF%2Bg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 2083', but common name is 'GVTC SHAKEN Cert 2083' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Nov 23 20:21 UTC