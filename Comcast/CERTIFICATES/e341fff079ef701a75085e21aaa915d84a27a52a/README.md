# STIR/SHAKEN CA Ecosystem Compliance
## Comcast

### Certificate e341fff079ef701a75085e21aaa915d84a27a52a
Tested At: 2022-10-28 18:15:09 +0000 UTC\
Initial Validity Period: 7300 day(s)\
Remaining Validity Period: 6345 day(s)\
Subject: CN=Comcast SHAKEN Root CA, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Root CA, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US

Link: 

View: [Click to view](https://understandingwebpki.com/?cert=MIICNzCCAdygAwIBAgIJAOKHGyEWYfhOMAoGCCqGSM49BAMCMG4xCzAJBgNVBAYTAlVTMRUwEwYDVQQIDAxQZW5uc3lsdmFuaWExFTATBgNVBAcMDFBoaWxhZGVscGhpYTEQMA4GA1UECgwHQ29tY2FzdDEfMB0GA1UEAwwWQ29tY2FzdCBTSEFLRU4gUm9vdCBDQTAeFw0yMDAzMTcxNDQ1MTVaFw00MDAzMTIxNDQ1MTVaMG4xCzAJBgNVBAYTAlVTMRUwEwYDVQQIDAxQZW5uc3lsdmFuaWExFTATBgNVBAcMDFBoaWxhZGVscGhpYTEQMA4GA1UECgwHQ29tY2FzdDEfMB0GA1UEAwwWQ29tY2FzdCBTSEFLRU4gUm9vdCBDQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJSzZzTeNhxovLtywhdzQU109JZLkcxXjBV9XB%2Fjgfv9qy9ZJfcP7x9crjSbzBu1%2BIoG65Qgvg5FGz5W6XR1cKKjYzBhMB0GA1UdDgQWBBSRkMqxhg5PFl6%2BtTdRP2l55SMbHDAfBgNVHSMEGDAWgBSRkMqxhg5PFl6%2BtTdRP2l55SMbHDAPBgNVHRMBAf8EBTADAQH%2FMA4GA1UdDwEB%2FwQEAwIBhjAKBggqhkjOPQQDAgNJADBGAiEAgfeTR2ucIDYaJvRyWW3VBUyhbAWVrAOQ7ZhPQpJ%2BcY8CIQDYYvj3djJDGthqwmbJRyxjX3YpxuHsoHR8plf7bHN5dA%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| n_pki_ca_key_usage | notice | SHAKEN PKI Best Practice | For CA certificates, the Key Usage extension should contain a single key usage value of keyCertSign |
| w_pki_ca_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### Not Effective

- e_cp1_3_ca_key_usage_crl_sign
- e_sti_ca_version
- e_sti_ca_serial_number
- e_sti_ca_subject_cn
- e_sti_root_authority_key_identifier
- e_sti_ca_key_usage
- e_sti_ca_signature_algorithm
- e_sti_ca_issuer_dn
- e_sti_ca_subject_public_key
- e_sti_root_extension_unknown
- e_sti_basic_constraints
- e_sti_root_certificate_policies
- e_sti_ca_subject
- e_sti_ca_subject_key_identifier

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 28/10/2022 at 18:15:47