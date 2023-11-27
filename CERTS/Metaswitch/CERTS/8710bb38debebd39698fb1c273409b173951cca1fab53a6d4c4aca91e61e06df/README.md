# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Everstream SHAKEN Cert 472C 

Tested At: 27 Nov 23 22:21 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 174 day(s)\
Subject: CN=Everstream SHAKEN Cert 472C\\ , O=Everstream, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/5dc472732f3dd53854a5183f1dadceeb6127fac7

[View certificate details](https://understandingwebpki.com/?cert=MIICTzCCAfagAwIBAgIQcw0F3VCn8B659gFZxPargzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDUyMDIyMjAyMFoXDTI0MDUxOTIyMjAyMFowSTELMAkGA1UEBhMCVVMxEzARBgNVBAoMCkV2ZXJzdHJlYW0xJTAjBgNVBAMMHEV2ZXJzdHJlYW0gU0hBS0VOIENlcnQgNDcyQyAwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATPE4X9mAFCgZElNKjjG%2Bnl79D4ZmPBRw20%2BmugqL4iFU%2BJpajcE%2FjyQCpthyXw3Mdd76hr1LCZyPxOYvx6e8GFo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ0NzJDMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFOCUts3fEgRfLulkY3U4LtrsFzO0MB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0cAMEQCIBxfNiqRj6JhNN%2FeEqDWp1uEX0qZZ1ENQYpqD%2FndP%2BzdAiA2b91n%2Bjb7tiqGGtyx%2BbDbJUyOmrNDIfQyi9R7AeBAyQ%3D%3D)

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


Generated: 27 Nov 23 22:56 UTC