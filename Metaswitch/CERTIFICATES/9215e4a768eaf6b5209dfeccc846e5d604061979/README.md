# STIR/SHAKEN CA Ecosystem Compliance
## Metaswitch

### Certificate 9215e4a768eaf6b5209dfeccc846e5d604061979
Tested At: 2022-10-27 22:43:40 +0000 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 521 day(s)\
Subject: CN=USCellular SHAKEN Cert 6349, L=Ch, ST=IL, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1

Link: https://crs.sti.uscellular.com/certs/uscc_shaken_CA.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIICVDCCAfqgAwIBAgIQZxVLbPqJw5UINoYY2jqr1jAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDQwMTE1MjgwOFoXDTI0MDMzMTE1MjgwOFowTTELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAklMMQswCQYDVQQHDAJDaDEkMCIGA1UEAwwbVVNDZWxsdWxhciBTSEFLRU4gQ2VydCA2MzQ5MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEiEoISdkvzCjey078Gz%2FzNVnHeJvthSEgrK95WIXUG4yVRm%2FubkEiInh53VZlW27VHM02mWLmf5eIn8QOxmDtM6OB2zCB2DAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIF4DAWBggrBgEFBQcBGgQKMAigBhYENjM0OTBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBT1FRvpTRjqLKOVkjVOuFFf1kU3ojAfBgNVHSMEGDAWgBTNHqcAEBDaMh1pGjnV0kYLLDyH1jAKBggqhkjOPQQDAgNIADBFAiEA%2BT8JXAe03JuPfBApV5psimLfX6WhkB3U2FRr0ubt7soCICZ8VWUSCjILY5ZS5sa3JZoIyh%2FiUPC2H0SO1UQPTC8p)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_key_usage | error | ATIS-1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| e_sti_issuer_dn | error | ATIS-1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| e_sti_subject | error | ATIS-1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### Not Effective

- e_sti_serial_number
- e_sti_signature_algorithm
- e_cp1_3_subject_sn
- e_sti_extension_unknown
- e_sti_subject_cn
- e_cp1_3_ambiguous_identifier
- w_cp_1_3_subject_email

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 27/10/2022 at 22:44:50