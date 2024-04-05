# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Mediacom 846F

Tested At: 05 Apr 24 18:56 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 51 day(s)\
Subject: CN=Mediacom 846F, O=Mediacom Communications Corporation, L=Chester, ST=New York, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://shaken.stir.mediacomcc.com/certs/mediacomcertchain.crt

[View certificate details](https://x509.io/?cert=MIICfjCCAiWgAwIBAgIQcT4%2BqcBbeC%2BEwJIU%2Ft5suTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDUyNzEzMjgxN1oXDTI0MDUyNjEzMjgxN1oweDELMAkGA1UEBhMCVVMxETAPBgNVBAgMCE5ldyBZb3JrMRAwDgYDVQQHDAdDaGVzdGVyMSwwKgYDVQQKDCNNZWRpYWNvbSBDb21tdW5pY2F0aW9ucyBDb3Jwb3JhdGlvbjEWMBQGA1UEAwwNTWVkaWFjb20gODQ2RjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABK1fxtkYqXgEF6LuRSxOoqfSwoYRTksJln6lGpSon7lUjNyCzvCDq1TEs0rFW57CObZvRBB0Ba9suFfv9EoOf26jgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDg0NkYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU28gepGTZ0%2BSNB45bqa6vg4mz8d4wHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDRwAwRAIgJ%2F4sn%2BN8M8OcdNmfv9hGkQYXtl5V0Gc4fcg8bIe72HMCIFWwKDWjO81ccw19MabBgS6tcD7l9YXJS6sUdbsNcfto)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'Mediacom 846F' does not contain 'SHAKEN' |

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


Generated: 05 Apr 24 19:04 UTC