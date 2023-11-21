# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Summit Broadband SHAKEN Cert 7857

Tested At: 21 Nov 23 01:22 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 832 day(s)\
Subject: CN=Summit Broadband SHAKEN Cert 7857, O=Summit Broadband, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/47c6d4e4af62db63cec43e632b0c057700d9704d

[View certificate details](https://understandingwebpki.com/?cert=MIICWjCCAgGgAwIBAgIQJNPP%2B%2FySec22mLpAX6rGajAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIzMDMwMjE3NTc1NloXDTI2MDMwMTE3NTc1NlowVDELMAkGA1UEBhMCVVMxGTAXBgNVBAoMEFN1bW1pdCBCcm9hZGJhbmQxKjAoBgNVBAMMIVN1bW1pdCBCcm9hZGJhbmQgU0hBS0VOIENlcnQgNzg1NzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABMyLqXdhK8EeOHfMuIyitpdGUh%2B8KPo5wBTTcckEOldIezCFViQh%2Fy88IO3snKfACJOvTy9HOUsHo1iB4stgC%2FujgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDc4NTcwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUq%2BI7o4sgMIvI7CGHFo8Shk12SOIwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDRwAwRAIgcbGpqWhu2AOAi8%2BY4vFqPUqlqIgpetzdflOorEZOgUkCIA2qU5uZJvo0Ly9GUKqZayLrVVqjvu5gXrLRFHrs5f7i)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 7857' |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 01:55 UTC