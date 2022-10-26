# STIR/SHAKEN CA Ecosystem Compliance
## Ribbon Communications

### Certificate 5e00d019ee3884ec7d8513deb8f4e8671a1db86e
Tested At: 2022-10-26 22:29:40 +0000 UTC\
Initial Validity Period: 9131 day(s)\
Remaining Validity Period: 8600 day(s)\
Subject: CN=SHAKEN Ribbon Root CA, OU=Certification Authority, O=Ribbon Communications, C=US\
Issuer: CN=SHAKEN Ribbon Root CA, OU=Certification Authority, O=Ribbon Communications, C=US

Link: 

View: [Click to view](https://understandingwebpki.com/?cert=MIICHzCCAcSgAwIBAgIQdbMkHP%2FQn78dmgdxeLv0djAKBggqhkjOPQQDAjBvMQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSAwHgYDVQQLExdDZXJ0aWZpY2F0aW9uIEF1dGhvcml0eTEeMBwGA1UEAxMVU0hBS0VOIFJpYmJvbiBSb290IENBMB4XDTIxMDUxMzAwMDAwMFoXDTQ2MDUxMjIzNTk1OVowbzELMAkGA1UEBhMCVVMxHjAcBgNVBAoTFVJpYmJvbiBDb21tdW5pY2F0aW9uczEgMB4GA1UECxMXQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkxHjAcBgNVBAMTFVNIQUtFTiBSaWJib24gUm9vdCBDQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABCY6GnVTWMs4FS83ne8vp9hz6F1Vk%2BthRDGpSb1jaLa8ZWhUtVbcUfoE%2F1bw3kspDyjuG0%2BjkJUaKqEmQkYnDOSjQjBAMA8GA1UdEwEB%2FwQFMAMBAf8wDgYDVR0PAQH%2FBAQDAgEGMB0GA1UdDgQWBBS7bhaSFwYNkmN%2BwZQyNeVu0dUUZDAKBggqhkjOPQQDAgNJADBGAiEAucCJi5Vvv0xHWAr6GXzyiZqG%2FbiL%2Bcrcr1JHVCmMbxICIQC8wkpOWO4DJpA3N54ig2xCp5dsoiibF4BJx64gYCZLww%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| n_pki_ca_key_usage | notice | SHAKEN PKI Best Practice | For CA certificates, the Key Usage extension should contain a single key usage value of keyCertSign |

150 tests were ran and no warning or error level issues were found

### Not Effective

- e_sti_ca_serial_number
- e_sti_ca_subject_public_key
- e_sti_ca_signature_algorithm
- e_sti_ca_subject
- e_sti_ca_key_usage
- e_sti_root_certificate_policies
- e_cp1_3_ca_key_usage_crl_sign
- e_sti_basic_constraints
- e_sti_ca_issuer_dn
- e_sti_ca_version
- w_cp1_3_ca_subject_rdn_unknown
- e_sti_ca_subject_key_identifier
- e_sti_ca_subject_cn
- e_sti_root_extension_unknown

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 26/10/2022 at 22:31:35