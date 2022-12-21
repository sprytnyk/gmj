package main

var Emojis = map[string]map[string]string{
	"art": {
		"value": "🎨",
		"usage": "Improve structure / format of the code.",
	},
	"zap": {
		"value": "⚡️",
		"usage": "Improve performance.",
	},
	"fire": {
		"value": "🔥",
		"usage": "Remove code or files.",
	},
	"bug": {
		"value": "🐛",
		"usage": "Fix a bug.",
	},
	"ambulance": {
		"value": "🚑️",
		"usage": "Critical hotfix.",
	},
	"sparkles": {
		"value": "✨",
		"usage": "Introduce new features.",
	},
	"memo": {
		"value": "📝",
		"usage": "Add or update documentation.",
	},
	"rocket": {
		"value": "🚀",
		"usage": "Deploy stuff.",
	},
	"lipstick": {
		"value": "💄",
		"usage": "Add or update the UI and style files.",
	},
	"tada": {
		"value": "🎉",
		"usage": "Begin a project.",
	},
	"white_check_mark": {
		"value": "✅",
		"usage": "Add, update, or pass tests.",
	},
	"lock": {
		"value": "🔒️",
		"usage": "Fix security issues.",
	},
	"closed_lock_with_key": {
		"value": "🔐",
		"usage": "Add or update secrets.",
	},
	"bookmark": {
		"value": "🔖",
		"usage": "Release / Version tags.",
	},
	"rotating_light": {
		"value": "🚨",
		"usage": "Fix compiler / linter warnings.",
	},
	"construction": {
		"value": "🚧",
		"usage": "Work in progress.",
	},
	"green_heart": {
		"value": "💚",
		"usage": "Fix CI Build.",
	},
	"arrow_down": {
		"value": "⬇️",
		"usage": "Downgrade dependencies.",
	},
	"arrow_up": {
		"value": "⬆️",
		"usage": "Upgrade dependencies.",
	},
	"pushpin": {
		"value": "📌",
		"usage": "Pin dependencies to specific versions.",
	},
	"construction_worker": {
		"value": "👷",
		"usage": "Add or update CI build system.",
	},
	"chart_with_upwards_trend": {
		"value": "📈",
		"usage": "Add or update analytics or track code.",
	},
	"recycle": {
		"value": "♻️",
		"usage": "Refactor code.",
	},
	"heavy_plus_sign": {
		"value": "➕",
		"usage": "Add a dependency.",
	},
	"heavy_minus_sign": {
		"value": "➖",
		"usage": "Remove a dependency.",
	},
	"wrench": {
		"value": "🔧",
		"usage": "Add or update configuration files.",
	},
	"hammer": {
		"value": "🔨",
		"usage": "Add or update development scripts.",
	},
	"globe_with_meridians": {
		"value": "🌐",
		"usage": "Internationalization and localization.",
	},
	"pencil2": {
		"value": "✏️",
		"usage": "Fix typos.",
	},
	"poop": {
		"value": "💩",
		"usage": "Write bad code that needs to be improved.",
	},
	"rewind": {
		"value": "⏪️",
		"usage": "Revert changes.",
	},
	"twisted_rightwards_arrows": {
		"value": "🔀",
		"usage": "Merge branches.",
	},
	"package": {
		"value": "📦️",
		"usage": "Add or update compiled files or packages.",
	},
	"alien": {
		"value": "👽️",
		"usage": "Update code due to external API changes.",
	},
	"truck": {
		"value": "🚚",
		"usage": "Move or rename resources (e.g.: files, paths, routes).",
	},
	"page_facing_up": {
		"value": "📄",
		"usage": "Add or update license.",
	},
	"boom": {
		"value": "💥",
		"usage": "Introduce breaking changes.",
	},
	"bento": {
		"value": "🍱",
		"usage": "Add or update assets.",
	},
	"wheelchair": {
		"value": "♿️",
		"usage": "Improve accessibility.",
	},
	"bulb": {
		"value": "💡",
		"usage": "Add or update comments in source code.",
	},
	"beers": {
		"value": "🍻",
		"usage": "Write code drunkenly.",
	},
	"speech_balloon": {
		"value": "💬",
		"usage": "Add or update text and literals.",
	},
	"card_file_box": {
		"value": "🗃️",
		"usage": "Perform database related changes.",
	},
	"loud_sound": {
		"value": "🔊",
		"usage": "Add or update logs.",
	},
	"mute": {
		"value": "🔇",
		"usage": "Remove logs.",
	},
	"busts_in_silhouette": {
		"value": "👥",
		"usage": "Add or update contributor(s).",
	},
	"children_crossing": {
		"value": "🚸",
		"usage": "Improve user experience / usability.",
	},
	"building_construction": {
		"value": "🏗️",
		"usage": "Make architectural changes.",
	},
	"iphone": {
		"value": "📱",
		"usage": "Work on responsive design.",
	},
	"clown_face": {
		"value": "🤡",
		"usage": "Mock things.",
	},
	"egg": {
		"value": "🥚",
		"usage": "Add or update an easter egg.",
	},
	"see_no_evil": {
		"value": "🙈",
		"usage": "Add or update a .gitignore file.",
	},
	"camera_flash": {
		"value": "📸",
		"usage": "Add or update snapshots.",
	},
	"alembic": {
		"value": "⚗️",
		"usage": "Perform experiments.",
	},
	"mag": {
		"value": "🔍️",
		"usage": "Improve SEO.",
	},
	"label": {
		"value": "🏷️",
		"usage": "Add or update types.",
	},
	"seedling": {
		"value": "🌱",
		"usage": "Add or update seed files.",
	},
	"triangular_flag_on_post": {
		"value": "🚩",
		"usage": "Add, update, or remove feature flags.",
	},
	"goal_net": {
		"value": "🥅",
		"usage": "Catch errors.",
	},
	"dizzy": {
		"value": "💫",
		"usage": "Add or update animations and transitions.",
	},
	"wastebasket": {
		"value": "🗑️",
		"usage": "Deprecate code that needs to be cleaned up.",
	},
	"passport_control": {
		"value": "🛂",
		"usage": "Work on code related to authorization, roles and permissions.",
	},
	"adhesive_bandage": {
		"value": "🩹",
		"usage": "Simple fix for a non-critical issue.",
	},
	"monocle_face": {
		"value": "🧐",
		"usage": "Data exploration/inspection.",
	},
	"coffin": {
		"value": "⚰️",
		"usage": "Remove dead code.",
	},
	"test_tube": {
		"value": "🧪",
		"usage": "Add a failing test.",
	},
	"necktie": {
		"value": "👔",
		"usage": "Add or update business logic.",
	},
	"stethoscope": {
		"value": "🩺",
		"usage": "Add or update healthcheck.",
	},
	"bricks": {
		"value": "🧱",
		"usage": "Infrastructure related changes.",
	},
	"technologist": {
		"value": "🧑",
		"usage": "Improve developer experience.",
	},
	"money_with_wings": {
		"value": "💸",
		"usage": "Add sponsorships or money related infrastructure.",
	},
	"thread": {
		"value": "🧵",
		"usage": "Add or update code related to multithreading or concurrency.",
	},
	"safety_vest": {
		"value": "🦺",
		"usage": "Add or update code related to validation.",
	},
}
