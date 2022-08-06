(window.webpackJsonp=window.webpackJsonp||[]).push([[117],{681:function(e,t,s){"use strict";s.r(t);var i=s(1),n=Object(i.a)({},(function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[s("h1",{attrs:{id:"adr-051-double-signing-risk-reduction"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#adr-051-double-signing-risk-reduction"}},[e._v("#")]),e._v(" ADR 051: Double Signing Risk Reduction")]),e._v(" "),s("h2",{attrs:{id:"changelog"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#changelog"}},[e._v("#")]),e._v(" Changelog")]),e._v(" "),s("ul",[s("li",[e._v("27-11-2019: Initial draft")]),e._v(" "),s("li",[e._v("13-01-2020: Separate into 2 ADR, This ADR will only cover Double signing Protection and ADR-052 handle Tendermint Mode")]),e._v(" "),s("li",[e._v('22-01-2020: change the title from "Double signing Protection" to "Double Signing Risk Reduction"')])]),e._v(" "),s("h2",{attrs:{id:"context"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#context"}},[e._v("#")]),e._v(" Context")]),e._v(" "),s("p",[e._v("To provide a risk reduction method for double signing incidents mistakenly executed by validators")]),e._v(" "),s("ul",[s("li",[e._v("Validators often mistakenly run duplicated validators to cause double-signing incident")]),e._v(" "),s("li",[e._v("This proposed feature is to reduce the risk of mistaken double-signing incident by checking recent N blocks before voting begins")]),e._v(" "),s("li",[e._v("When we think of such serious impact on double-signing incident, it is very reasonable to have multiple risk reduction algorithm built in node daemon")])]),e._v(" "),s("h2",{attrs:{id:"decision"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#decision"}},[e._v("#")]),e._v(" Decision")]),e._v(" "),s("p",[e._v("We would like to suggest a double signing risk reduction method.")]),e._v(" "),s("ul",[s("li",[e._v("Methodology : query recent consensus results to find out whether node's consensus key is used on consensus recently or not")]),e._v(" "),s("li",[e._v("When to check\n"),s("ul",[s("li",[e._v("When the state machine starts "),s("code",[e._v("ConsensusReactor")]),e._v(" after fully synced")]),e._v(" "),s("li",[e._v("When the node is validator ( with privValidator )")]),e._v(" "),s("li",[e._v("When "),s("code",[e._v("cs.config.DoubleSignCheckHeight > 0")])])])]),e._v(" "),s("li",[e._v("How to check\n"),s("ol",[s("li",[e._v("When a validator is transformed from syncing status to fully synced status, the state machine check recent N blocks ("),s("code",[e._v("latest_height - double_sign_check_height")]),e._v(") to find out whether there exists consensus votes using the validator's consensus key")]),e._v(" "),s("li",[e._v("If there exists votes from the validator's consensus key, exit state machine program")])])]),e._v(" "),s("li",[e._v("Configuration\n"),s("ul",[s("li",[e._v("We would like to suggest by introducing "),s("code",[e._v("double_sign_check_height")]),e._v(" parameter in "),s("code",[e._v("config.toml")]),e._v(" and cli, how many blocks state machine looks back to check votes")]),e._v(" "),s("li",[s("span",{pre:!0},[s("code",[e._v("double_sign_check_height = {{ .Consensus.DoubleSignCheckHeight }}")])]),e._v(" in "),s("code",[e._v("config.toml")])]),e._v(" "),s("li",[s("code",[e._v("tendermint node --consensus.double_sign_check_height")]),e._v(" in cli")]),e._v(" "),s("li",[e._v("State machine ignore checking procedure when "),s("code",[e._v("double_sign_check_height == 0")])])])])]),e._v(" "),s("h2",{attrs:{id:"status"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#status"}},[e._v("#")]),e._v(" Status")]),e._v(" "),s("p",[e._v("Implemented")]),e._v(" "),s("h2",{attrs:{id:"consequences"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#consequences"}},[e._v("#")]),e._v(" Consequences")]),e._v(" "),s("h3",{attrs:{id:"positive"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#positive"}},[e._v("#")]),e._v(" Positive")]),e._v(" "),s("ul",[s("li",[e._v("Validators can avoid double signing incident by mistakes. (eg. If another validator node is voting on consensus, starting new validator node with same consensus key will cause panic stop of the state machine because consensus votes with the consensus key are found in recent blocks)")]),e._v(" "),s("li",[e._v("We expect this method will prevent majority of double signing incident by mistakes.")])]),e._v(" "),s("h3",{attrs:{id:"negative"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#negative"}},[e._v("#")]),e._v(" Negative")]),e._v(" "),s("ul",[s("li",[e._v("When the risk reduction method is on, restarting a validator node will panic because the node itself voted on consensus with the same consensus key. So, validators should stop the state machine, wait for some blocks, and then restart the state machine to avoid panic stop.")])]),e._v(" "),s("h3",{attrs:{id:"neutral"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#neutral"}},[e._v("#")]),e._v(" Neutral")]),e._v(" "),s("h2",{attrs:{id:"references"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#references"}},[e._v("#")]),e._v(" References")]),e._v(" "),s("ul",[s("li",[e._v("Issue "),s("a",{attrs:{href:"https://github.com/tendermint/tendermint/issues/4059",target:"_blank",rel:"noopener noreferrer"}},[e._v("#4059"),s("OutboundLink")],1),e._v(" : double-signing protection")])])])}),[],!1,null,null,null);t.default=n.exports}}]);