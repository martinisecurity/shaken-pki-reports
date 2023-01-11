# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Kaplan Telephone SHAKEN cert 0432

Tested At: 11 Jan 23 21:04 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 844 day(s)\
Subject: CN=Kaplan Telephone SHAKEN cert 0432, O=Kaplan Telephone, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/ef7499777a884734baaa98de26c0caed26fe57be

[View certificate details](https://understandingwebpki.com/?cert=MIICWjCCAgGgAwIBAgIQbvTJmJG8Gb2AKsf6BdBwnzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIyMDUwNTE0MTQ0OFoXDTI1MDUwNDE0MTQ0OFowVDELMAkGA1UEBhMCVVMxGTAXBgNVBAoMEEthcGxhbiBUZWxlcGhvbmUxKjAoBgNVBAMMIUthcGxhbiBUZWxlcGhvbmUgU0hBS0VOIGNlcnQgMDQzMjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABCE9vDw81Ftfpf5oSKPGSMFDJ%2F7QVT1LDE0nHXrsk9lCDYnG3RKRQjyqw9srkLa9kg6dgLY7efM4bwyrDXdIwQWjgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDA0MzIwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQU1o%2B2W%2B1mVn3QO1eWdG8d1dyQVjQwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDRwAwRAIgFCVnamac8UgJxQb5LJbhZPl0gry3YfMVAQ1NmNbOPe8CIH2PJBFO33Qw%2B8qaUxpDFj7oqyKXT3Dnxu86naJW%2BIur)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 0432' |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 11 Jan 23 21:04 UTC