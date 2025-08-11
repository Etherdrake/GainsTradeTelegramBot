package topicmaps

var EventnameToTopicHashTrading = map[string]string{
	"Done":                     "0xc3a6f986261de9467c2838c6df8ef74f9107855159205600c0bc7a14cdfd3888",
	"Paused":                   "0x0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd2",
	"NumberUpdated":            "0x8cf3e35f6221b16e1670a3413180c9484bf5aa71787905909fa82a6a2662e9ab",
	"BypassTriggerLinkUpdated": "0x3c0f648064e21cceb91d918a80bed80a58c69474a28d50fe3d9975b72be97c44",
	"MarketOrderInitiated":     "0x3e544118c04e3bb18b669475695cd270ba0e41fb13177483f01c14222de62a86",
	"OpenLimitPlaced":          "0xdfabd6f206f17b7f2e1f9e0d33c40d30d1e8d7b6a4f520a03fdc1c1811059343",
	"OpenLimitUpdated":         "0x710a8db87f04e82a9de40076812593a965f4aa48693196d2144c07ff9710e890",
	"OpenLimitCanceled":        "0xf1b38881d7f4b2b12141c5f39c5124545d6112532eb6afbe9630cdbde3ee53e9",
	"TpUpdated":                "0x7e06a81c7a47891ccc7455b5ccb2ed850e32bb655ccda67eb3ebaaeed83242a4",
	"SlUpdated":                "0x1fc4a6c7ffe506697979b8ed54dc4135cd1ecd26a2745f70b760a2492222b316",
	"NftOrderInitiated":        "0x50a583b02839381dff332433f1a37825291992d796b87483d7c51649ef504d43",
	"ChainlinkCallbackTimeout": "0x3adaa586cdbe84dd24e45bd7dada6da933d7c2d1c7b4e4cd02fce033356decb1",
	"CouldNotCloseTrade":       "0x60e497734ddabcd7293fd91739aaf65cf525eb539c97be528125a235a89288d8",
}

var TopicHashToEventnameTrading = map[string]string{
	"0xc3a6f986261de9467c2838c6df8ef74f9107855159205600c0bc7a14cdfd3888": "Done",
	"0x0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd2": "Paused",
	"0x8cf3e35f6221b16e1670a3413180c9484bf5aa71787905909fa82a6a2662e9ab": "NumberUpdated",
	"0x3c0f648064e21cceb91d918a80bed80a58c69474a28d50fe3d9975b72be97c44": "BypassTriggerLinkUpdated",
	"0x3e544118c04e3bb18b669475695cd270ba0e41fb13177483f01c14222de62a86": "MarketOrderInitiated", // DONE DIAMOND
	"0xdfabd6f206f17b7f2e1f9e0d33c40d30d1e8d7b6a4f520a03fdc1c1811059343": "OpenLimitPlaced",      // DONE DIAMOND
	"0x710a8db87f04e82a9de40076812593a965f4aa48693196d2144c07ff9710e890": "OpenLimitUpdated",
	"0xf1b38881d7f4b2b12141c5f39c5124545d6112532eb6afbe9630cdbde3ee53e9": "OpenLimitCanceled",
	"0x7e06a81c7a47891ccc7455b5ccb2ed850e32bb655ccda67eb3ebaaeed83242a4": "TpUpdated", // DONE DIAMOND
	"0x1fc4a6c7ffe506697979b8ed54dc4135cd1ecd26a2745f70b760a2492222b316": "SlUpdated",
	"0x50a583b02839381dff332433f1a37825291992d796b87483d7c51649ef504d43": "NftOrderInitiated", // DONE DIAMOND
	"0x3adaa586cdbe84dd24e45bd7dada6da933d7c2d1c7b4e4cd02fce033356decb1": "ChainlinkCallbackTimeout",
	"0x60e497734ddabcd7293fd91739aaf65cf525eb539c97be528125a235a89288d8": "CouldNotCloseTrade",
}

var EventnameToTopichashCallbacks = map[string]string{
	"MarketExecuted":           "0xca42b0e44cd853d207b87e8f8914eaefef9c9463a8c77ca33754aa62f6904f00", // DONE DIAMOND
	"LimitExecuted":            "0xa97091b8c54bf9d1906c2a06322d0ea74fedde4538cdcdf95d81d0ffdca41857", // DONE DIAMOND
	"MarketOpenCanceled":       "0x1dc3532663e5566091476fb5aba1e514ef733714c83d4feec5723de6f16c3269",
	"MarketCloseCanceled":      "0x293df767d6749666902026d2f6a2cc4e5f15cdede46402226c42ef4fdf27a17c",
	"NftOrderCanceled":         "0xe9681b5336d843735c62e93114e5a0f45912a84ae83fa3f3ed80ca5ad933dfc3",
	"ClosingFeeSharesPUpdated": "0x0caa98ed9a1605da290817d1f67b1b83c63f9229abeca5123df5d90581c49558",
	"Pause":                    "0x9422424b175dda897495a07b091ef74a3ef715cf6d866fc972954c1c7f459304",
	"Done":                     "0xc3a6f986261de9467c2838c6df8ef74f9107855159205600c0bc7a14cdfd3888",
	"GovFeesClaimed":           "0x39b06677afbdb5c3b9934c9ce55728be609a055c398ecd957e8d19a5d3d80a5b",
	"GovFeeCharged":            "0xccd80d359a6fbe0bfa5cbb1ecf0854adbe8c67b4ed6bf10d3c0d78c2be0f48cb",
	"ReferralFeeCharged":       "0x0f5273269f52308b9c40fafda3ca13cc42f715fcd795365e87f351f59e249313",
	"TriggerFeeCharged":        "0x17fa86cf4833d28c6224a940e6bd001f2db0cb3d89d69727765679b3efee6559",
	"SssFeeCharged":            "0xd1e388cc27c5125a80cf538c12b26dc5a784071d324a81a736e4d17f238588e4", // DONE DIAMOND
	"DaiVaultFeeCharged":       "0x60c73da98faf96842eabd77a0c73964cd189dbaf2c9ae90923a3fed137f30e3e", // DONE DIAMOND
	"BorrowingFeeCharged":      "0xe7d34775bf6fd7b34e703a903ef79ab16166ebdffce96a66f4d2f84b6263bb29",
	"PairMaxLeverageUpdated":   "0x95924bc10431f9a625a06fe5a27d55f4348510b2da42a18fe3bf2a6f2c4eab67",
}

var TopichashToEventnameCallbacks = map[string]string{
	"0xca42b0e44cd853d207b87e8f8914eaefef9c9463a8c77ca33754aa62f6904f00": "MarketExecuted", // DONE DIAMOND
	"0xa97091b8c54bf9d1906c2a06322d0ea74fedde4538cdcdf95d81d0ffdca41857": "LimitExecuted",  // DONE DIAMOND
	"0x1dc3532663e5566091476fb5aba1e514ef733714c83d4feec5723de6f16c3269": "MarketOpenCanceled",
	"0x293df767d6749666902026d2f6a2cc4e5f15cdede46402226c42ef4fdf27a17c": "MarketCloseCanceled",
	"0xe9681b5336d843735c62e93114e5a0f45912a84ae83fa3f3ed80ca5ad933dfc3": "NftOrderCanceled",
	"0x0caa98ed9a1605da290817d1f67b1b83c63f9229abeca5123df5d90581c49558": "ClosingFeeSharesPUpdated",
	"0x9422424b175dda897495a07b091ef74a3ef715cf6d866fc972954c1c7f459304": "Pause",
	"0xc3a6f986261de9467c2838c6df8ef74f9107855159205600c0bc7a14cdfd3888": "Done",
	"0x39b06677afbdb5c3b9934c9ce55728be609a055c398ecd957e8d19a5d3d80a5b": "GovFeesClaimed",
	"0xccd80d359a6fbe0bfa5cbb1ecf0854adbe8c67b4ed6bf10d3c0d78c2be0f48cb": "GovFeeCharged",
	"0x0f5273269f52308b9c40fafda3ca13cc42f715fcd795365e87f351f59e249313": "ReferralFeeCharged",
	"0x17fa86cf4833d28c6224a940e6bd001f2db0cb3d89d69727765679b3efee6559": "TriggerFeeCharged",
	"0xd1e388cc27c5125a80cf538c12b26dc5a784071d324a81a736e4d17f238588e4": "SssFeeCharged",      // DONE DIAMOND
	"0x60c73da98faf96842eabd77a0c73964cd189dbaf2c9ae90923a3fed137f30e3e": "DaiVaultFeeCharged", // DONE DIAMOND
	"0xe7d34775bf6fd7b34e703a903ef79ab16166ebdffce96a66f4d2f84b6263bb29": "BorrowingFeeCharged",
	"0x95924bc10431f9a625a06fe5a27d55f4348510b2da42a18fe3bf2a6f2c4eab67": "PairMaxLeverageUpdated",
}
