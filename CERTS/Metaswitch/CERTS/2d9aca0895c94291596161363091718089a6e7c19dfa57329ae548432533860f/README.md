# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Telesystem SHAKEN Cert 786E

Tested At: 31 Oct 22 16:42 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 454 day(s)\
Subject: CN=Telesystem SHAKEN Cert 786E, O=Telesystem, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/1cb5524539bfa481384ff17da116fd8eed39952f

View: [Click to view](https://understandingwebpki.com/?cert=MIICTjCCAfWgAwIBAgIQNlHLcp4gkfS9wy8zuWTXjzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDEyNzE3MTYzM1oXDTI0MDEyNzE3MTYzM1owSDELMAkGA1UEBhMCVVMxEzARBgNVBAoMClRlbGVzeXN0ZW0xJDAiBgNVBAMMG1RlbGVzeXN0ZW0gU0hBS0VOIENlcnQgNzg2RTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABNGUjtU4%2FpQ0l9VP6d7u9GO6%2BUT7FXmuw4HYP0vsiAST3MZu21LEiyf0zVqnN7G7tTfJ3qX9SY6Sf1Zw7Tfkke6jgdswgdgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCBeAwFgYIKwYBBQUHARoECjAIoAYWBDc4NkUwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUpQvuDCQ0Z8KrH%2Bo99etj1k1dODMwHwYDVR0jBBgwFoAUzR6nABAQ2jIdaRo51dJGCyw8h9YwCgYIKoZIzj0EAwIDRwAwRAIgKqZJIbso4%2Bt6PHBAr06ShA58CxbvNUkX8vX%2FC%2F1qfAkCIChs4rTNFej6CyulDMv%2FmtmBUZ7xi1mP1tMB5cg5qmbC)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_sti_issuer_dn](../../ISSUES/e_sti_issuer_dn/README.md) | error | ATIS&#x2011;1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| [e_sti_key_usage](../../ISSUES/e_sti_key_usage/README.md) | error | ATIS&#x2011;1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |


### Not Effective

- e_cp1_3_ambiguous_identifier
- e_cp1_3_subject_sn
- e_sti_extension_unknown
- e_sti_serial_number
- e_sti_signature_algorithm
- e_sti_subject_cn
- w_cp_1_3_subject_email

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 31/10/2022 at 16:43:22