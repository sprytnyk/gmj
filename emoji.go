package main

// Emojis map contains all available codepoints to be used in commit messages.
var Emojis = map[string]map[string]string{
	"art": {
		"codepoint": "🎨",
		"usage":     "Improve structure / format of the code.",
	},
	"zap": {
		"codepoint": "⚡️",
		"usage":     "Improve performance.",
	},
	"fire": {
		"codepoint": "🔥",
		"usage":     "Remove code or files.",
	},
	"bug": {
		"codepoint": "🐛",
		"usage":     "Fix a bug.",
	},
	"ambulance": {
		"codepoint": "🚑️",
		"usage":     "Critical hotfix.",
	},
	"sparkles": {
		"codepoint": "✨",
		"usage":     "Introduce new features.",
	},
	"memo": {
		"codepoint": "📝",
		"usage":     "Add or update documentation.",
	},
	"rocket": {
		"codepoint": "🚀",
		"usage":     "Deploy stuff.",
	},
	"lipstick": {
		"codepoint": "💄",
		"usage":     "Add or update the UI and style files.",
	},
	"tada": {
		"codepoint": "🎉",
		"usage":     "Begin a project.",
	},
	"white_check_mark": {
		"codepoint": "✅",
		"usage":     "Add, update, or pass tests.",
	},
	"lock": {
		"codepoint": "🔒️",
		"usage":     "Fix security issues.",
	},
	"closed_lock_with_key": {
		"codepoint": "🔐",
		"usage":     "Add or update secrets.",
	},
	"bookmark": {
		"codepoint": "🔖",
		"usage":     "Release / Version tags.",
	},
	"rotating_light": {
		"codepoint": "🚨",
		"usage":     "Fix compiler / linter warnings.",
	},
	"construction": {
		"codepoint": "🚧",
		"usage":     "Work in progress.",
	},
	"green_heart": {
		"codepoint": "💚",
		"usage":     "Fix CI Build.",
	},
	"arrow_down": {
		"codepoint": "⬇️",
		"usage":     "Downgrade dependencies.",
	},
	"arrow_up": {
		"codepoint": "⬆️",
		"usage":     "Upgrade dependencies.",
	},
	"pushpin": {
		"codepoint": "📌",
		"usage":     "Pin dependencies to specific versions.",
	},
	"construction_worker": {
		"codepoint": "👷",
		"usage":     "Add or update CI build system.",
	},
	"chart_with_upwards_trend": {
		"codepoint": "📈",
		"usage":     "Add or update analytics or track code.",
	},
	"recycle": {
		"codepoint": "♻️",
		"usage":     "Refactor code.",
	},
	"heavy_plus_sign": {
		"codepoint": "➕",
		"usage":     "Add a dependency.",
	},
	"heavy_minus_sign": {
		"codepoint": "➖",
		"usage":     "Remove a dependency.",
	},
	"wrench": {
		"codepoint": "🔧",
		"usage":     "Add or update configuration files.",
	},
	"hammer": {
		"codepoint": "🔨",
		"usage":     "Add or update development scripts.",
	},
	"globe_with_meridians": {
		"codepoint": "🌐",
		"usage":     "Internationalization and localization.",
	},
	"pencil2": {
		"codepoint": "✏️",
		"usage":     "Fix typos.",
	},
	"poop": {
		"codepoint": "💩",
		"usage":     "Write bad code that needs to be improved.",
	},
	"rewind": {
		"codepoint": "⏪️",
		"usage":     "Revert changes.",
	},
	"twisted_rightwards_arrows": {
		"codepoint": "🔀",
		"usage":     "Merge branches.",
	},
	"package": {
		"codepoint": "📦️",
		"usage":     "Add or update compiled files or packages.",
	},
	"alien": {
		"codepoint": "👽️",
		"usage":     "Update code due to external API changes.",
	},
	"truck": {
		"codepoint": "🚚",
		"usage":     "Move or rename resources (e.g.: files, paths, routes).",
	},
	"page_facing_up": {
		"codepoint": "📄",
		"usage":     "Add or update license.",
	},
	"boom": {
		"codepoint": "💥",
		"usage":     "Introduce breaking changes.",
	},
	"bento": {
		"codepoint": "🍱",
		"usage":     "Add or update assets.",
	},
	"wheelchair": {
		"codepoint": "♿️",
		"usage":     "Improve accessibility.",
	},
	"bulb": {
		"codepoint": "💡",
		"usage":     "Add or update comments in source code.",
	},
	"beers": {
		"codepoint": "🍻",
		"usage":     "Write code drunkenly.",
	},
	"speech_balloon": {
		"codepoint": "💬",
		"usage":     "Add or update text and literals.",
	},
	"card_file_box": {
		"codepoint": "🗃️",
		"usage":     "Perform database related changes.",
	},
	"loud_sound": {
		"codepoint": "🔊",
		"usage":     "Add or update logs.",
	},
	"mute": {
		"codepoint": "🔇",
		"usage":     "Remove logs.",
	},
	"busts_in_silhouette": {
		"codepoint": "👥",
		"usage":     "Add or update contributor(s).",
	},
	"children_crossing": {
		"codepoint": "🚸",
		"usage":     "Improve user experience / usability.",
	},
	"building_construction": {
		"codepoint": "🏗️",
		"usage":     "Make architectural changes.",
	},
	"iphone": {
		"codepoint": "📱",
		"usage":     "Work on responsive design.",
	},
	"clown_face": {
		"codepoint": "🤡",
		"usage":     "Mock things.",
	},
	"egg": {
		"codepoint": "🥚",
		"usage":     "Add or update an easter egg.",
	},
	"see_no_evil": {
		"codepoint": "🙈",
		"usage":     "Add or update a .gitignore file.",
	},
	"camera_flash": {
		"codepoint": "📸",
		"usage":     "Add or update snapshots.",
	},
	"alembic": {
		"codepoint": "⚗️",
		"usage":     "Perform experiments.",
	},
	"mag": {
		"codepoint": "🔍️",
		"usage":     "Improve SEO.",
	},
	"label": {
		"codepoint": "🏷️",
		"usage":     "Add or update types.",
	},
	"seedling": {
		"codepoint": "🌱",
		"usage":     "Add or update seed files.",
	},
	"triangular_flag_on_post": {
		"codepoint": "🚩",
		"usage":     "Add, update, or remove feature flags.",
	},
	"goal_net": {
		"codepoint": "🥅",
		"usage":     "Catch errors.",
	},
	"dizzy": {
		"codepoint": "💫",
		"usage":     "Add or update animations and transitions.",
	},
	"wastebasket": {
		"codepoint": "🗑️",
		"usage":     "Deprecate code that needs to be cleaned up.",
	},
	"passport_control": {
		"codepoint": "🛂",
		"usage":     "Work on code related to authorization, roles and permissions.",
	},
	"adhesive_bandage": {
		"codepoint": "🩹",
		"usage":     "Simple fix for a non-critical issue.",
	},
	"monocle_face": {
		"codepoint": "🧐",
		"usage":     "Data exploration/inspection.",
	},
	"coffin": {
		"codepoint": "⚰️",
		"usage":     "Remove dead code.",
	},
	"test_tube": {
		"codepoint": "🧪",
		"usage":     "Add a failing test.",
	},
	"necktie": {
		"codepoint": "👔",
		"usage":     "Add or update business logic.",
	},
	"stethoscope": {
		"codepoint": "🩺",
		"usage":     "Add or update healthcheck.",
	},
	"bricks": {
		"codepoint": "🧱",
		"usage":     "Infrastructure related changes.",
	},
	"technologist": {
		"codepoint": "🧑‍💻",
		"usage":     "Improve developer experience.",
	},
	"money_with_wings": {
		"codepoint": "💸",
		"usage":     "Add sponsorships or money related infrastructure.",
	},
	"thread": {
		"codepoint": "🧵",
		"usage":     "Add or update code related to multithreading or concurrency.",
	},
	"safety_vest": {
		"codepoint": "🦺",
		"usage":     "Add or update code related to validation.",
	},
}
