# STIR/SHAKEN CA Ecosystem Compliance

## Telonium

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 39 certificates were included in the corpus being tested
- 6 certificates in the corpus were skipped because they are duplicates
- 12 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 21 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 23.81% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 472 days is the average remaining validity for the certificates in the corpus
- 480 days is the average initial validity for the certificates in the corpus
- 2 certificates expire in the next 30 days
- 1.00 average number of unexpired certificates per OCN observed
- 21 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 2 | [e_atis_ext_crl_distribution_struct](ISSUES/e_atis_ext_crl_distribution_struct/README.md) | ATIS1000080 |
| 3 | [e_atis_ext_not_specified](ISSUES/e_atis_ext_not_specified/README.md) | ATIS1000080 |

#### CA Certificates

- 5 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 5 certificates being tested against the remaining rules
- 4.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 60.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 4255 days is the average remaining validity for the certificates in the corpus
- 4309 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_ext_certificate_policies_ca](ISSUES/e_atis_ext_certificate_policies_ca/README.md) | ATIS1000080 |
| 1 | [e_atis_ext_crl_distribution_ca](ISSUES/e_atis_ext_crl_distribution_ca/README.md) | ATIS1000080 |
| 2 | [e_atis_ext_key_usage](ISSUES/e_atis_ext_key_usage/README.md) | ATIS1000080 |
| 1 | [e_atis_subject_c_iso_ca](ISSUES/e_atis_subject_c_iso_ca/README.md) | ATIS1000080 |
| 3 | [e_atis_subject_cn_ca](ISSUES/e_atis_subject_cn_ca/README.md) | ATIS1000080 |
| 2 | [e_atis_subject_cn_root](ISSUES/e_atis_subject_cn_root/README.md) | ATIS1000080 |
| 1 | [e_atis_subject_dn_ca](ISSUES/e_atis_subject_dn_ca/README.md) | ATIS1000080 |
| 1 | [e_shaken_certificate_policies_id_ca](ISSUES/e_shaken_certificate_policies_id_ca/README.md) | US_SHAKEN_CP |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 10&#160;Jul&#160;23&#160;14:44&#160;UTC | SHAKEN 656K | 28&#160;Jun&#160;25&#160;19:35&#160;UTC | true | [view](CERTS/fd285a02878cd6c995f532bbc9b51b4721f5dd3828aa271194760254eb1cd549/README.md) |
| 11&#160;Jul&#160;23&#160;18:15&#160;UTC | SHAKEN 159K | 22&#160;Jun&#160;25&#160;22:42&#160;UTC | true | [view](CERTS/d9b2af8d51b09c75e70a5c6d60de90a663be51c44f141639708ab7a6cf151dfa/README.md) |
| 11&#160;Jul&#160;23&#160;18:45&#160;UTC | SHAKEN 421K | 21&#160;Jun&#160;25&#160;04:00&#160;UTC | true | [view](CERTS/cfaa04fe3747df8c203aa057940d1d7b1b455b15c8da2ba73cf7384797edfad8/README.md) |
| 01&#160;Aug&#160;23&#160;17:17&#160;UTC | SHAKEN 143K | 31&#160;Jul&#160;25&#160;17:18&#160;UTC | false | [view](CERTS/7888e1431475305a982a18f6722eca01f235f420b9968f2f30f49937dbb6ed00/README.md) |
| 25&#160;Aug&#160;23&#160;20:58&#160;UTC | SHAKEN 709K | 25&#160;Aug&#160;24&#160;20:59&#160;UTC | false | [view](CERTS/50f6d51678c38fc547c53e5e649ae55172323da164c4fa76036ede9290544532/README.md) |
| 08&#160;Sep&#160;23&#160;23:10&#160;UTC | SHAKEN 721K | 08&#160;Sep&#160;24&#160;23:11&#160;UTC | false | [view](CERTS/a3988074881dad0e96799028c8f312ec3dd145810a538e92747761cc0cc344a9/README.md) |
| 09&#160;Jan&#160;24&#160;14:50&#160;UTC | SHAKEN 785K | 09&#160;Jan&#160;25&#160;14:51&#160;UTC | false | [view](CERTS/06142e30bf65c757ac31cbc4950162477d131f67f9086019b17a6fa24773be0f/README.md) |
| 30&#160;Jan&#160;24&#160;20:38&#160;UTC | SHAKEN 442K | 30&#160;Jan&#160;25&#160;20:39&#160;UTC | false | [view](CERTS/4ba867185817379be2eb96c2f1048418362641680814f06ca86338c96b89d51c/README.md) |
| 07&#160;Feb&#160;24&#160;17:04&#160;UTC | SHAKEN 830K | 07&#160;Feb&#160;25&#160;17:05&#160;UTC | false | [view](CERTS/0d2d8f2b9fa8f4b8f35133e87be3bee44b51cc55318af9ef230a6db9f2981eec/README.md) |
| 20&#160;Feb&#160;24&#160;15:17&#160;UTC | SHAKEN 824K | 20&#160;Feb&#160;25&#160;15:18&#160;UTC | false | [view](CERTS/637d69144054e0a7e304433f3a9a1fec16fe102732c19945cf5f43f214b7faf4/README.md) |
| 27&#160;Feb&#160;24&#160;16:46&#160;UTC | SHAKEN 963J | 27&#160;Feb&#160;26&#160;16:47&#160;UTC | true | [view](CERTS/ceef9f6d88a1efc4c056ecaf0f5d42c5f106fb5fc895c8ccc1f7ea97b6ac1093/README.md) |
| 08&#160;Mar&#160;24&#160;19:11&#160;UTC | SHAKEN 854K | 08&#160;Mar&#160;25&#160;19:12&#160;UTC | true | [view](CERTS/bfea21afc2db20c52f74b16c054d12ff6e839acc8b18401aced0154ef7e03750/README.md) |
| 12&#160;Mar&#160;24&#160;18:22&#160;UTC | SHAKEN 869K | 12&#160;Mar&#160;25&#160;18:23&#160;UTC | false | [view](CERTS/c373220caf192639c39402b0527c7817d54b588cc151d2da504962cec365c99d/README.md) |
| 22&#160;Apr&#160;24&#160;18:38&#160;UTC | SHAKEN 698K | 22&#160;Apr&#160;25&#160;18:39&#160;UTC | false | [view](CERTS/745f739166e82d8e3974488040c8ec3c25e05bfa3f02f8f6f631ad8580256f5b/README.md) |
| 26&#160;Apr&#160;24&#160;15:16&#160;UTC | SHAKEN 856K | 26&#160;Apr&#160;25&#160;15:17&#160;UTC | false | [view](CERTS/5e39e1ae5bf58085df8e42422e4ade4981dc97992b8f3752060b82202fa7280c/README.md) |
| 30&#160;Apr&#160;24&#160;15:58&#160;UTC | SHAKEN 902K | 30&#160;Apr&#160;25&#160;15:59&#160;UTC | false | [view](CERTS/de7f61878b7336e75166e84320fcbb0ebce13837c032a691680c71d71ac933f0/README.md) |
| 17&#160;May&#160;24&#160;16:06&#160;UTC | SHAKEN 622K | 15&#160;Jun&#160;25&#160;18:34&#160;UTC | false | [view](CERTS/8910d8a2c6bcd7dea8f26e0fab11c8d5b4e3360ce849533605347c71dd427918/README.md) |
| 20&#160;May&#160;24&#160;13:49&#160;UTC | SHAKEN 865K | 20&#160;May&#160;25&#160;13:50&#160;UTC | false | [view](CERTS/386e494b04018ba0d7cee2d8dd5d5b39da967d93c0d83160e1674a81237afd2d/README.md) |
| 20&#160;Jun&#160;24&#160;19:05&#160;UTC | SHAKEN 633K | 27&#160;Dec&#160;25&#160;16:31&#160;UTC | false | [view](CERTS/5c0d32d70b614b08cec3502a7ef34663daa44eb814894c021def70d56c113d62/README.md) |
| 26&#160;Jun&#160;24&#160;21:01&#160;UTC | SHAKEN 930K | 26&#160;Jun&#160;25&#160;21:02&#160;UTC | false | [view](CERTS/9cadceccf69a1eb0ad9c47404f506cb771df0ec254c87e9f1c7d66b1cb680eee/README.md) |
| 12&#160;Aug&#160;24&#160;16:23&#160;UTC | SHAKEN 715K | 11&#160;Sep&#160;26&#160;18:30&#160;UTC | false | [view](CERTS/396ffc900acd18b268a476122bfaae0835afda50c8648220980093051eb2bbfb/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 08&#160;Mar&#160;23&#160;18:40&#160;UTC | Telonium STI-CA Root1 | 05&#160;Mar&#160;35&#160;18:40&#160;UTC | true | [view](CERTS/96c66865ce5558c2ce3723c0b414538fcacadcd0f3286108fef57dc447f122f9/README.md) |
| 08&#160;Mar&#160;23&#160;18:40&#160;UTC | Telonium STI-CA Root2 | 07&#160;Mar&#160;38&#160;18:40&#160;UTC | true | [view](CERTS/a58b27999411d3d54121d4eadc82aa128be1fef96cda3029b2015677188ea40b/README.md) |
| 09&#160;Mar&#160;23&#160;15:18&#160;UTC | Telonium STI-CA Intermediate CA | 06&#160;Mar&#160;33&#160;15:18&#160;UTC | true | [view](CERTS/7c701216e591c9a3b84550ff46566dd420c7f182eb3cfc5abe5739cdbe271169/README.md) |
| 21&#160;Jul&#160;23&#160;00:49&#160;UTC | Telonium SHAKEN ROOT G1 | 21&#160;Jul&#160;35&#160;00:49&#160;UTC | false | [view](CERTS/37e1a126fc5d84ff59f332b2fe8196205bd0e4f7353be497ad17770d9ca6cea5/README.md) |
| 01&#160;Aug&#160;23&#160;16:36&#160;UTC | Telonium SHAKEN Intermediate G1 | 01&#160;Aug&#160;33&#160;16:36&#160;UTC | false | [view](CERTS/8c0ef8682826bec79a8c64881899f6a5a4a1d52dfebe28ae419c23f85df96ea0/README.md) |


Generated: 22 Aug 24 15:44 UTC