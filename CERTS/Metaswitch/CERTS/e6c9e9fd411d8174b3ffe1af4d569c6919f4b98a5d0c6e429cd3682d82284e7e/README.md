# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Mediacom 846F

Tested At: 31 Oct 22 18:17 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 573 day(s)\
Subject: CN=Mediacom 846F, O=Mediacom Communications Corporation, L=Chester, ST=New York, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://shaken.stir.mediacomcc.com/certs/mediacomcertchain.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIICfjCCAiWgAwIBAgIQcT4%2BqcBbeC%2BEwJIU%2Ft5suTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDUyNzEzMjgxN1oXDTI0MDUyNjEzMjgxN1oweDELMAkGA1UEBhMCVVMxETAPBgNVBAgMCE5ldyBZb3JrMRAwDgYDVQQHDAdDaGVzdGVyMSwwKgYDVQQKDCNNZWRpYWNvbSBDb21tdW5pY2F0aW9ucyBDb3Jwb3JhdGlvbjEWMBQGA1UEAwwNTWVkaWFjb20gODQ2RjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABK1fxtkYqXgEF6LuRSxOoqfSwoYRTksJln6lGpSon7lUjNyCzvCDq1TEs0rFW57CObZvRBB0Ba9suFfv9EoOf26jgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDg0NkYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU28gepGTZ0%2BSNB45bqa6vg4mz8d4wHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDRwAwRAIgJ%2F4sn%2BN8M8OcdNmfv9hGkQYXtl5V0Gc4fcg8bIe72HMCIFWwKDWjO81ccw19MabBgS6tcD7l9YXJS6sUdbsNcfto)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_sti_issuer_dn](../../ISSUES/e_sti_issuer_dn/README.md) | error | ATIS&#x2011;1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_sti_key_usage](../../ISSUES/e_sti_key_usage/README.md) | error | ATIS&#x2011;1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


### Not Effective

- e_cp1_3_ambiguous_identifier
- e_cp1_3_subject_sn
- e_sti_extension_unknown
- e_sti_serial_number
- e_sti_signature_algorithm
- e_sti_subject_cn
- w_cp_1_3_subject_email

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 31/10/2022 at 18:25:03