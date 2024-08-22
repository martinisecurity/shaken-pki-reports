# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Nemont SHAKEN Cert 2247

Tested At: 22 Aug 24 15:16 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: -94 day(s)\
Subject: CN=Nemont SHAKEN Cert 2247, O=Nemont, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/da7847867ee00785d849b849e374c81ecd3f2545

[View certificate details](https://x509.io/?cert=MIICRjCCAe2gAwIBAgIQfnlUqz%2FjGhWuQpkRz9hSCDAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDUyMDIyMTMzNFoXDTI0MDUxOTIyMTMzNFowQDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBk5lbW9udDEgMB4GA1UEAwwXTmVtb250IFNIQUtFTiBDZXJ0IDIyNDcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAS%2BrXdXOqbl9jRCjhcjBTMbvnQSnqlU9G22qqqbWZjbrXdhMPPR7wiKayplO0VXvJ0JJvDBWCPozGRvvTAWBqt1o4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQyMjQ3MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFHEAsNJ6HPhOXQifjP39H1IxFXnqMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0cAMEQCIGOm5sWFQpVJRWQkydwwLEul7TNdnekR8c8CaWNVBatkAiBltUHKtrRAKeR%2FEIWrzkCnxkVyUd3V%2BHYZEwW3i3cufw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |

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


Generated: 22 Aug 24 15:44 UTC