# STIR/SHAKEN CA Ecosystem Compliance
## Metaswitch

### Certificate 2bed8a9246c78dfdf1cbac12bd20316ad5892163
Tested At: 2022-10-27 22:31:51 +0000 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 506 day(s)\
Subject: CN=RCN SHAKEN Cert 7615, O=RCN, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1

Link: https://sti-cr.cgah.tnsi.com/certs/11d9e0e06eeaaf56fb9be54eec6aae5a8b3461fe

View: [Click to view](https://understandingwebpki.com/?cert=MIICQTCCAeegAwIBAgIQCFfTK6Nnq8%2BDNtDYu%2BZYnzAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDMxNzE3MDYwMVoXDTI0MDMxNjE3MDYwMVowOjELMAkGA1UEBhMCVVMxDDAKBgNVBAoMA1JDTjEdMBsGA1UEAwwUUkNOIFNIQUtFTiBDZXJ0IDc2MTUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATAt6ELjr2XfUM%2FjKVUUjVMaZcM45EZ%2FuxWR4jx0mAVh8Q6Uvfrxq8h8gnbLVtqhV%2FRMz0VNdsrqrG0WUpG4X3po4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQ3NjE1MEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFHd6jgfsvzke%2Bwfzg1VrzP25KaPeMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0gAMEUCICUcEjuJCsxEFC6zuTL7lSoVrTdYAyCDB%2Be8qnO4bol1AiEAtX1wCQcGbJmq4L9Hnu2WmV5xO1m4ccoKQHdKu52oYh0%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_issuer_dn | error | ATIS-1000080 | The DN shall contain a Country (C=) attribute, a Common Name (CN=) attribute and an Organization (O=) attribute |
| e_sti_key_usage | error | ATIS-1000080 | The Key Usage extension shall contain a single key usage value of digitalSignature |

### Not Effective

- w_cp_1_3_subject_email
- e_cp1_3_subject_sn
- e_sti_extension_unknown
- w_cp1_3_subject_rdn_unknown
- e_sti_serial_number
- e_sti_subject_cn
- e_sti_signature_algorithm
- e_cp1_3_ambiguous_identifier

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 27/10/2022 at 22:33:03