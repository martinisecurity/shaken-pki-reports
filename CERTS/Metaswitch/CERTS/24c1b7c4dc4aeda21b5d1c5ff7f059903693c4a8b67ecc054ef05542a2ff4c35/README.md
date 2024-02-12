# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Kaplan Telephone SHAKEN cert 0432

Tested At: 12 Feb 24 16:26 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 447 day(s)\
Subject: CN=Kaplan Telephone SHAKEN cert 0432, O=Kaplan Telephone, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/c57d0a622b6d80dae91a7391b608bcf4c3d88f74

[View certificate details](https://understandingwebpki.com/?cert=MIICWjCCAgGgAwIBAgIQbvTJmJG8Gb2AKsf6BdBwnzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIyMDUwNTE0MTQ0OFoXDTI1MDUwNDE0MTQ0OFowVDELMAkGA1UEBhMCVVMxGTAXBgNVBAoMEEthcGxhbiBUZWxlcGhvbmUxKjAoBgNVBAMMIUthcGxhbiBUZWxlcGhvbmUgU0hBS0VOIGNlcnQgMDQzMjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABCE9vDw81Ftfpf5oSKPGSMFDJ%2F7QVT1LDE0nHXrsk9lCDYnG3RKRQjyqw9srkLa9kg6dgLY7efM4bwyrDXdIwQWjgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDA0MzIwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQU1o%2B2W%2B1mVn3QO1eWdG8d1dyQVjQwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDRwAwRAIgFCVnamac8UgJxQb5LJbhZPl0gry3YfMVAQ1NmNbOPe8CIH2PJBFO33Qw%2B8qaUxpDFj7oqyKXT3Dnxu86naJW%2BIur)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 0432', but common name is 'Kaplan Telephone SHAKEN cert 0432' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 12 Feb 24 17:02 UTC