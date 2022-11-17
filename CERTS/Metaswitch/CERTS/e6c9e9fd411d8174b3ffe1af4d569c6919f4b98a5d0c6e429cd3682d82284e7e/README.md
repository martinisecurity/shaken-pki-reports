# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Mediacom 846F

Tested At: 17 Nov 22 19:12 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 556 day(s)\
Subject: CN=Mediacom 846F, O=Mediacom Communications Corporation, L=Chester, ST=New York, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://shaken.stir.mediacomcc.com/certs/mediacomcertchain.crt

[View certificate details](https://understandingwebpki.com/?cert=MIICfjCCAiWgAwIBAgIQcT4%2BqcBbeC%2BEwJIU%2Ft5suTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDUyNzEzMjgxN1oXDTI0MDUyNjEzMjgxN1oweDELMAkGA1UEBhMCVVMxETAPBgNVBAgMCE5ldyBZb3JrMRAwDgYDVQQHDAdDaGVzdGVyMSwwKgYDVQQKDCNNZWRpYWNvbSBDb21tdW5pY2F0aW9ucyBDb3Jwb3JhdGlvbjEWMBQGA1UEAwwNTWVkaWFjb20gODQ2RjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABK1fxtkYqXgEF6LuRSxOoqfSwoYRTksJln6lGpSon7lUjNyCzvCDq1TEs0rFW57CObZvRBB0Ba9suFfv9EoOf26jgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDg0NkYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU28gepGTZ0%2BSNB45bqa6vg4mz8d4wHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDRwAwRAIgJ%2F4sn%2BN8M8OcdNmfv9hGkQYXtl5V0Gc4fcg8bIe72HMCIFWwKDWjO81ccw19MabBgS6tcD7l9YXJS6sUdbsNcfto)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_issuer_dn](../../ISSUES/e_atis_issuer_dn/README.md) | error | ATIS1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 17 Nov 22 19:20 UTC