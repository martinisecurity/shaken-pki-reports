# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Leaf Certificates

- Average validity span as of leaf certificates 67 days
- Percentage of leaf certificates expiring in the next 30 days is 79.49%
- Certificates with Errors: 78
- Certificates with Warnings: 78
- Certificates with Notices: 0
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 2

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| error | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md#leaf-certificates) | RFC5280 | 75 |
| error | [e_sti_certificate_policies](ISSUES/e_sti_certificate_policies/README.md#leaf-certificates) | ATIS-1000080 | 3 |
| not effective | [w_cp_1_3_subject_email](ISSUES/w_cp_1_3_subject_email/README.md#leaf-certificates) | United States SHAKEN CP | 1 |
| error | [e_cp1_3_subject_sn](ISSUES/e_cp1_3_subject_sn/README.md#leaf-certificates) | United States SHAKEN CP | 78 |
| warn | [w_pki_subject_rdn_unknown](ISSUES/w_pki_subject_rdn_unknown/README.md#leaf-certificates) | SHAKEN PKI Best Practice | 78 |
| error | [e_cp1_3_ambiguous_identifier](ISSUES/e_cp1_3_ambiguous_identifier/README.md#leaf-certificates) | United States SHAKEN CP | 4 |
| not effective | [e_sti_subject_cn](ISSUES/e_sti_subject_cn/README.md#leaf-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_extension_unknown](ISSUES/e_sti_extension_unknown/README.md#leaf-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_signature_algorithm](ISSUES/e_sti_signature_algorithm/README.md#leaf-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_serial_number](ISSUES/e_sti_serial_number/README.md#leaf-certificates) | ATIS-1000080 | 2 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 238 certificates skipped because they are currently expired.\
\*\*\* 1 certificates skipped because they are not currently trusted by the STI-PA.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 02 Nov 21 16:40 UTC | Charter Communications Inc SHAKEN 5606 | true | [view](CERTIFICATES/abbd6bca303f79a930d3a14e0d3538c51934a97a/README.md) |
| 13 Jan 22 14:17 UTC | MobileSphere SHAKEN 873J | true | [view](CERTIFICATES/701a8780085449f2f6ee50c287cecc20f2219b2f/README.md) |
| 30 Mar 22 12:54 UTC | Fusion Connect SHAKEN 2720 | true | [view](CERTIFICATES/136f6839b2fa440965940ff380d9cde7b053f2d0/README.md) |
| 27 Apr 22 18:22 UTC | CCI SHAKEN 663J | true | [view](CERTIFICATES/7c0be1458e20b8247bce9338ecbe8322578f807c/README.md) |
| 10 Jun 22 14:00 UTC | SHAKEN 597J | true | [view](CERTIFICATES/9e8c6e83efe9133df13009a69d1dfbe8801f0a13/README.md) |
| 20 Jun 22 20:25 UTC | SHAKEN 857J | true | [view](CERTIFICATES/caf6b73e79063f79bf34dae854ebf059f32c1a45/README.md) |
| 22 Jun 22 18:46 UTC | SHAKEN 755J | true | [view](CERTIFICATES/ee364dffdc58cce8eb91842e65530ece378acec1/README.md) |
| 29 Jun 22 20:24 UTC | SHAKEN 736J | true | [view](CERTIFICATES/8a1091afcbec190e8877ecdd666ef18c50cb9356/README.md) |
| 25 Jul 22 18:36 UTC | SHAKEN 578J | true | [view](CERTIFICATES/b0affbea73591da4097b4915d92b53818409f67f/README.md) |
| 10 Aug 22 18:11 UTC | SHAKEN 073H | true | [view](CERTIFICATES/0ae828173955498af37d944351d38f36446db6d6/README.md) |
| 01 Sep 22 03:25 UTC | SHAKEN 6628 | true | [view](CERTIFICATES/01505cd524c060dffca87d0b043e0be889026713/README.md) |
| 01 Sep 22 12:31 UTC | SHAKEN 193E | true | [view](CERTIFICATES/b6d72ada27949e72427f323a03e4ddc7e6d2e803/README.md) |
| 13 Sep 22 17:24 UTC | SHAKEN 706J | true | [view](CERTIFICATES/8c58c1679cf840f228b7d5285f431a25ad4d6562/README.md) |
| 26 Sep 22 18:40 UTC | SHAKEN 505J | true | [view](CERTIFICATES/b110a3f4054ff7065b1c6f14883ade56f2ced44e/README.md) |
| 26 Sep 22 18:42 UTC | SHAKEN 505J | true | [view](CERTIFICATES/d3e15676cd711e2afcfebdfa1c0ea2186d50b5b9/README.md) |
| 30 Sep 22 22:34 UTC | SHAKEN 8526 | true | [view](CERTIFICATES/969863c8a4f335ed0a19906ae5c3558905b7725a/README.md) |
| 01 Oct 22 04:04 UTC | SHAKEN 345J | true | [view](CERTIFICATES/eb03b74587d3cfb43d588d8463e17c52db9244ed/README.md) |
| 01 Oct 22 04:34 UTC | SHAKEN 345J | true | [view](CERTIFICATES/22a8d66576c72f5acfc282c00a7f9e28863234ea/README.md) |
| 01 Oct 22 12:31 UTC | SHAKEN 193E | true | [view](CERTIFICATES/164675f5894b5459867b897576a59e6aba1fc656/README.md) |
| 01 Oct 22 15:23 UTC | SHAKEN 722J | true | [view](CERTIFICATES/006cd8457db666bfa0b1df0497d439dd83f9f246/README.md) |
| 10 Oct 22 18:30 UTC | SHAKEN 115K | true | [view](CERTIFICATES/dd733ba189741f0b583a3e595a3c14db8e87be01/README.md) |
| 10 Oct 22 18:39 UTC | SHAKEN 066K | true | [view](CERTIFICATES/b6bbcdc291fe28062d3ce9d6850e9b0c6fef05cc/README.md) |
| 12 Oct 22 19:39 UTC | SHAKEN 815G | true | [view](CERTIFICATES/e689269bc057ddf1876d610f8dd14761b439defc/README.md) |
| 13 Oct 22 18:02 UTC | SHAKEN 952J | true | [view](CERTIFICATES/414512f13c27bb40a6f47741eabd63eb5c90eed3/README.md) |
| 20 Oct 22 09:13 UTC | SHAKEN 841J | true | [view](CERTIFICATES/3889bee709733ff19f8acc2ae4a3761d288ffdfe/README.md) |
| 20 Oct 22 15:48 UTC | SHAKEN 622J | true | [view](CERTIFICATES/175e3b3de9a38a8355d047b901bd5eda0c1f1a6a/README.md) |
| 21 Oct 22 20:13 UTC | SHAKEN 983J | true | [view](CERTIFICATES/fa5b20f5d1c96d0095d20cd3cf723e84a876e7ab/README.md) |
| 22 Oct 22 00:21 UTC | SHAKEN 551G | true | [view](CERTIFICATES/fc9938ac7cc123cd2e27f38171de9b31183c5871/README.md) |
| 22 Oct 22 04:04 UTC | SHAKEN 345J | true | [view](CERTIFICATES/fc99f74650aae6ff90f41e781e215bf1a8f06f68/README.md) |
| 22 Oct 22 04:44 UTC | SHAKEN 345J | true | [view](CERTIFICATES/ecf5e6662e3084505582f48a356cc82f2e95a7dd/README.md) |
| 23 Oct 22 08:39 UTC | SHAKEN 012K | true | [view](CERTIFICATES/a96d0880d34d1cb60fd0f00ce985722fb952a02d/README.md) |
| 23 Oct 22 18:35 UTC | SHAKEN 089K | true | [view](CERTIFICATES/fde031a4737afd05ee9408293b200a35bfa307e4/README.md) |
| 23 Oct 22 18:37 UTC | SHAKEN 9714 | true | [view](CERTIFICATES/cb1430446ee425f297b477264f245934a792e826/README.md) |
| 23 Oct 22 19:56 UTC | SHAKEN 297K | true | [view](CERTIFICATES/9df80eed66106e1aa17a18c41d830a3e240c445a/README.md) |
| 23 Oct 22 20:17 UTC | SHAKEN 674J | true | [view](CERTIFICATES/ff10d41f263188f0218fa8d9f25b21d42e3fe1f7/README.md) |
| 23 Oct 22 20:18 UTC | SHAKEN 738J | true | [view](CERTIFICATES/bf14dbda90b2ffc796c34770c105c1ca625de56b/README.md) |
| 23 Oct 22 20:18 UTC | SHAKEN 7453 | true | [view](CERTIFICATES/e6d88e273d127842c50b03c4b5e1c0dd088872fe/README.md) |
| 23 Oct 22 20:21 UTC | SHAKEN 849J | true | [view](CERTIFICATES/ef4f17c2145aaea571f41b690925b622e8146800/README.md) |
| 23 Oct 22 20:21 UTC | SHAKEN 469A | true | [view](CERTIFICATES/b3972d4d4686a49e9ae279e328f8b822c476fe52/README.md) |
| 23 Oct 22 20:21 UTC | SHAKEN 790J | true | [view](CERTIFICATES/38fe760cf3e12683c131b3d0f2e891f007794503/README.md) |
| 23 Oct 22 20:22 UTC | SHAKEN 625J | true | [view](CERTIFICATES/e2aacd36d02951b3e73b552ea16b976f7ccf9f26/README.md) |
| 23 Oct 22 20:22 UTC | SHAKEN 738J | true | [view](CERTIFICATES/dcfb7c5055b8e4f07f3136116eaf7e354dabb82a/README.md) |
| 23 Oct 22 20:23 UTC | SHAKEN 459J | true | [view](CERTIFICATES/cf6536ed84fa4fb421ebf66f89eb4d51794270c3/README.md) |
| 23 Oct 22 20:24 UTC | SHAKEN 366G | true | [view](CERTIFICATES/bd7a8b9542e2a04e92af34f261cb8afa92a1af0d/README.md) |
| 23 Oct 22 20:24 UTC | SHAKEN 0226 | true | [view](CERTIFICATES/f9c467c1018ad1eabfa846fc2aa1c5c27cd8123f/README.md) |
| 23 Oct 22 20:24 UTC | SHAKEN 738J | true | [view](CERTIFICATES/e8f019de24ae5f9dd28301a9f73bb0ce59907a7f/README.md) |
| 23 Oct 22 20:48 UTC | SHAKEN 518J | true | [view](CERTIFICATES/76579174555c2405b844fbf6ace0ccb15e0aa02c/README.md) |
| 23 Oct 22 22:11 UTC | SHAKEN 0172 | true | [view](CERTIFICATES/8dcc3b92f9308b0c51305675d84fe107f0187ecc/README.md) |
| 24 Oct 22 16:32 UTC | SHAKEN 291K | true | [view](CERTIFICATES/2661ebd23f777002d5a8b8e0cced45b06bf12b10/README.md) |
| 24 Oct 22 20:14 UTC | SHAKEN 983J | true | [view](CERTIFICATES/17933d08d67709d3ad7021060216ceff05c5a63d/README.md) |
| 24 Oct 22 20:22 UTC | SHAKEN 738J | true | [view](CERTIFICATES/aa9b65dead966d217c2e51f1e78559a830fefbf8/README.md) |
| 24 Oct 22 20:51 UTC | SHAKEN 157C | true | [view](CERTIFICATES/21e6797a2a553ddf3e1d5c9d5ceb99f93fa6c773/README.md) |
| 24 Oct 22 21:41 UTC | SHAKEN 606F | true | [view](CERTIFICATES/0ea62e7be98fcc1d31d5d28965f8fabeb1cd9e5a/README.md) |
| 25 Oct 22 01:41 UTC | SHAKEN 278K | true | [view](CERTIFICATES/baa3a224db49a001f7036652f702def7906e41fc/README.md) |
| 25 Oct 22 14:33 UTC | SHAKEN 2550 | true | [view](CERTIFICATES/c612e0f95b5d7ac7c60edf495a4cbb05014f0c85/README.md) |
| 25 Oct 22 15:33 UTC | SHAKEN 722J | true | [view](CERTIFICATES/160aff282003da74ec5b92fd7f622fb8487e6879/README.md) |
| 26 Oct 22 08:39 UTC | SHAKEN 012K | true | [view](CERTIFICATES/d645bcf28ab9c9c17f50a509aee493520976cd68/README.md) |
| 26 Oct 22 10:43 UTC | SHAKEN 8526 | true | [view](CERTIFICATES/c75227976fe6d09fed2d445131b0aa220b68e55f/README.md) |
| 26 Oct 22 13:45 UTC | SHAKEN 056J | true | [view](CERTIFICATES/a28ecd097baf3964135df1ec46706a1b79cf49e9/README.md) |
| 26 Oct 22 17:24 UTC | SHAKEN 107K | true | [view](CERTIFICATES/3475a0a95ab1463443919df71ceff158ee753348/README.md) |
| 26 Oct 22 18:13 UTC | SHAKEN 060K | true | [view](CERTIFICATES/d312970974b165866f58d2ad927c368487d7219f/README.md) |
| 26 Oct 22 18:36 UTC | SHAKEN 089K | true | [view](CERTIFICATES/d33fcbf0b9ea652abf3b0deedaf23a6099b9fa64/README.md) |
| 26 Oct 22 18:38 UTC | SHAKEN 9714 | true | [view](CERTIFICATES/852e1a324d6698ca17de03508938b427cf1351ce/README.md) |
| 26 Oct 22 20:17 UTC | SHAKEN 674J | true | [view](CERTIFICATES/e209bd7e8c5e4fdede7214376441955ca4f9069c/README.md) |
| 26 Oct 22 20:18 UTC | SHAKEN 735J | true | [view](CERTIFICATES/f490353b428ee888e60e090869adc243627b15bd/README.md) |
| 26 Oct 22 20:19 UTC | SHAKEN 738J | true | [view](CERTIFICATES/49e1a5bd16f39e929c5141142996d289b1d34522/README.md) |
| 26 Oct 22 20:22 UTC | SHAKEN 469A | true | [view](CERTIFICATES/2a11d7d8e37ec0f245f516797c3815df9d6d6476/README.md) |
| 26 Oct 22 20:22 UTC | SHAKEN 790J | true | [view](CERTIFICATES/530153d7d36e6a903758e07a9b38a256b381df80/README.md) |
| 26 Oct 22 20:23 UTC | SHAKEN 625J | true | [view](CERTIFICATES/2c1fc9cdde6833c8f868003987fa38a40b1fc733/README.md) |
| 26 Oct 22 20:23 UTC | SHAKEN 738J | true | [view](CERTIFICATES/45954ff85e77c5fc6cbb2f406f86dc808fb27aa4/README.md) |
| 26 Oct 22 20:23 UTC | SHAKEN 459J | true | [view](CERTIFICATES/e4c0d5fe00cacd7db40b620d5dbf0fee81c2bb50/README.md) |
| 26 Oct 22 20:25 UTC | SHAKEN 366G | true | [view](CERTIFICATES/8e3532282d073bd1b3a55cc4062ce64023df766f/README.md) |
| 26 Oct 22 20:25 UTC | SHAKEN 738J | true | [view](CERTIFICATES/dce5a7287dd0c8ce81b627fea32ce6e7787519cf/README.md) |
| 26 Oct 22 20:25 UTC | SHAKEN 0226 | true | [view](CERTIFICATES/70975f08fe7ef842969e92792ff7516fd2bbc0ce/README.md) |
| 26 Oct 22 22:00 UTC | SHAKEN 551G | true | [view](CERTIFICATES/e4c1bd4a476e082df5a9198733d5a40f0b77d5c5/README.md) |
| 27 Oct 22 16:32 UTC | SHAKEN 291K | true | [view](CERTIFICATES/23f1ff180392009217e28b1ea8182175b16eea13/README.md) |
| 27 Oct 22 19:07 UTC | SHAKEN 749J | true | [view](CERTIFICATES/c2e542860a7f485376d3d6023f1436fb73ac0f3d/README.md) |
| 27 Oct 22 20:14 UTC | SHAKEN 983J | true | [view](CERTIFICATES/b52fcac4e6c751ea15757474af2a207d2626e3e1/README.md) |

\* For issues relating to this CAs certificate repositories see this [report](URL/README.md).

### CA Certificates

- Certificates with Errors: 3
- Certificates with Warnings: 3
- Certificates with Notices: 3
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 3

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| not effective | [e_sti_ca_subject_cn](ISSUES/e_sti_ca_subject_cn/README.md#ca-certificates) | ATIS-1000080 | 3 |
| not effective | [e_sti_ca_certificate_policies](ISSUES/e_sti_ca_certificate_policies/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_ca_crl_distribution](ISSUES/e_sti_ca_crl_distribution/README.md#ca-certificates) | ATIS-1000080 | 2 |
| error | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md#ca-certificates) | RFC5280 | 3 |
| notice | [n_pki_ca_key_usage](ISSUES/n_pki_ca_key_usage/README.md#ca-certificates) | SHAKEN PKI Best Practice | 3 |
| error | [e_cp1_3_ca_key_usage_crl_sign](ISSUES/e_cp1_3_ca_key_usage_crl_sign/README.md#ca-certificates) | United States SHAKEN CP | 3 |
| not effective | [n_sti_ca_certificate_policy_critical](ISSUES/n_sti_ca_certificate_policy_critical/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_root_certificate_policies](ISSUES/e_sti_root_certificate_policies/README.md#ca-certificates) | ATIS-1000080 | 1 |
| not effective | [e_sti_root_extension_unknown](ISSUES/e_sti_root_extension_unknown/README.md#ca-certificates) | ATIS-1000080 | 1 |
| warn | [w_pki_ca_subject_rdn_unknown](ISSUES/w_pki_ca_subject_rdn_unknown/README.md#ca-certificates) | SHAKEN PKI Best Practice | 3 |
| not effective | [e_sti_ca_serial_number](ISSUES/e_sti_ca_serial_number/README.md#ca-certificates) | ATIS-1000080 | 3 |
| not effective | [e_sti_ca_extension_unknown](ISSUES/e_sti_ca_extension_unknown/README.md#ca-certificates) | ATIS-1000080 | 2 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 0 certificates skipped because they are currently expired.\
\*\*\* 1 certificates skipped because they are not currently trusted by the STI-PA.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Root CA1 | true | [view](CERTIFICATES/36dc4ae1d521b8a5aedd10498e6ce757581b197f/README.md) |
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Issuing CA1 | true | [view](CERTIFICATES/2d540c7179aad19290a2098ccac8053645527154/README.md) |
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Issuing CA3 | true | [view](CERTIFICATES/df7d871ff60d213820a96308346eda870d6e8ed2/README.md) |

Generated: 28/10/2022 at 18:22:55