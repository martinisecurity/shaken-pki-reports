# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Everstream SHAKEN Cert 472C 

Tested At: 31 Oct 22 19:21 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 567 day(s)\
Subject: CN=Everstream SHAKEN Cert 472C\\ , O=Everstream, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://sti-cr.cgah.tnsi.com/certs/5dc472732f3dd53854a5183f1dadceeb6127fac7

View: [Click to view](https://understandingwebpki.com/?cert=MIICTzCCAfagAwIBAgIQcw0F3VCn8B659gFZxPargzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDUyMDIyMjAyMFoXDTI0MDUxOTIyMjAyMFowSTELMAkGA1UEBhMCVVMxEzARBgNVBAoMCkV2ZXJzdHJlYW0xJTAjBgNVBAMMHEV2ZXJzdHJlYW0gU0hBS0VOIENlcnQgNDcyQyAwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATPE4X9mAFCgZElNKjjG%2Bnl79D4ZmPBRw20%2BmugqL4iFU%2BJpajcE%2FjyQCpthyXw3Mdd76hr1LCZyPxOYvx6e8GFo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ0NzJDMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFOCUts3fEgRfLulkY3U4LtrsFzO0MB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0cAMEQCIBxfNiqRj6JhNN%2FeEqDWp1uEX0qZZ1ENQYpqD%2FndP%2BzdAiA2b91n%2Bjb7tiqGGtyx%2BbDbJUyOmrNDIfQyi9R7AeBAyQ%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_sti_key_usage](../../ISSUES/e_sti_key_usage/README.md) | error | ATIS&#x2011;1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |
| [e_sti_issuer_dn](../../ISSUES/e_sti_issuer_dn/README.md) | error | ATIS&#x2011;1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |


### Not Effective

- e_cp1_3_ambiguous_identifier
- e_cp1_3_subject_sn
- e_sti_extension_unknown
- e_sti_serial_number
- e_sti_signature_algorithm
- e_sti_subject_cn
- w_cp_1_3_subject_email

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 31/10/2022 at 19:21:49