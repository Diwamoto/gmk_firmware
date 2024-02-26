package tgk

const (
	MOD_KEY_MASK = 0xFF00
	TO_KEY_MASK  = 0x0010
)

const (
	TYPE_MEDIA = 0xE400
	TYPE_NOMAL = 0xF000
	TYPE_MOUSE = 0xD000
	TYPE_MOD   = 0xFF00
	TYPE_TO    = 0xFF10
)

const (
	JP_KC_LEFT  = TYPE_NOMAL | 0xE0
	JP_KC_RIGHT = TYPE_NOMAL | 0xE1
	JP_KC_LALT  = TYPE_NOMAL | 0xE2
	JP_KC_GUI   = TYPE_NOMAL | 0xE3
	JP_KC_RCTRL = TYPE_NOMAL | 0xE4
	JP_KC_RSHFT = TYPE_NOMAL | 0xE5

	JP_KC_BS = TYPE_NOMAL | 0x2A
)

const (
	KC_MOD1 = TYPE_MOD | 0x00
	KC_MOD2 = TYPE_MOD | 0x01
	KC_MOD3 = TYPE_MOD | 0x02
	KC_MOD4 = TYPE_MOD | 0x03
	KC_MOD5 = TYPE_MOD | 0x04
	KC_MOD6 = TYPE_MOD | 0x05

	KC_TO0 = TYPE_TO | 0x00
	KC_TO1 = TYPE_TO | 0x01
	KC_TO2 = TYPE_TO | 0x02
	KC_TO3 = TYPE_TO | 0x03
	KC_TO4 = TYPE_TO | 0x04
	KC_TO5 = TYPE_TO | 0x05
)

const (
	// restore default keymap for QMK
	KEY_RESTORE_DEFAULT_KEYMAP = 0x7C03
)

const (
	KC_A            = TYPE_NOMAL | 0x04
	KC_B            = TYPE_NOMAL | 0x05
	KC_C            = TYPE_NOMAL | 0x06
	KC_D            = TYPE_NOMAL | 0x07
	KC_E            = TYPE_NOMAL | 0x08
	KC_F            = TYPE_NOMAL | 0x09
	KC_G            = TYPE_NOMAL | 0x0A
	KC_H            = TYPE_NOMAL | 0x0B
	KC_I            = TYPE_NOMAL | 0x0C
	KC_J            = TYPE_NOMAL | 0x0D
	KC_K            = TYPE_NOMAL | 0x0E
	KC_L            = TYPE_NOMAL | 0x0F
	KC_M            = TYPE_NOMAL | 0x10
	KC_N            = TYPE_NOMAL | 0x11
	KC_O            = TYPE_NOMAL | 0x12
	KC_P            = TYPE_NOMAL | 0x13
	KC_Q            = TYPE_NOMAL | 0x14
	KC_R            = TYPE_NOMAL | 0x15
	KC_S            = TYPE_NOMAL | 0x16
	KC_T            = TYPE_NOMAL | 0x17
	KC_U            = TYPE_NOMAL | 0x18
	KC_V            = TYPE_NOMAL | 0x19
	KC_W            = TYPE_NOMAL | 0x1A
	KC_X            = TYPE_NOMAL | 0x1B
	KC_Y            = TYPE_NOMAL | 0x1C
	KC_Z            = TYPE_NOMAL | 0x1D
	KC_1            = TYPE_NOMAL | 0x1E
	KC_2            = TYPE_NOMAL | 0x1F
	KC_3            = TYPE_NOMAL | 0x20
	KC_4            = TYPE_NOMAL | 0x21
	KC_5            = TYPE_NOMAL | 0x22
	KC_6            = TYPE_NOMAL | 0x23
	KC_7            = TYPE_NOMAL | 0x24
	KC_8            = TYPE_NOMAL | 0x25
	KC_9            = TYPE_NOMAL | 0x26
	KC_0            = TYPE_NOMAL | 0x27
	KC_ENTER        = TYPE_NOMAL | 0x28
	KC_ESC          = TYPE_NOMAL | 0x29
	KC_BS           = TYPE_NOMAL | 0x2A
	KC_TAB          = TYPE_NOMAL | 0x2B
	KC_SPACE        = TYPE_NOMAL | 0x2C
	KC_MINUS        = TYPE_NOMAL | 0x2D
	KC_HAT          = TYPE_NOMAL | 0x2E
	KC_AT           = TYPE_NOMAL | 0x2F
	KC_LEFT_BRACE   = TYPE_NOMAL | 0x30
	KC_LBRACE       = KC_LEFT_BRACE
	KC_RIGHT_BRACE  = TYPE_NOMAL | 0x32
	KC_RBRACE       = KC_RIGHT_BRACE
	KC_SEMICOLON    = TYPE_NOMAL | 0x33
	KC_SEMI         = KC_SEMICOLON
	KC_COLON        = TYPE_NOMAL | 0x34
	KC_HANKAKU      = TYPE_NOMAL | 0x35
	KC_COMMA        = TYPE_NOMAL | 0x36
	KC_PERIOD       = TYPE_NOMAL | 0x37
	KC_SLASH        = TYPE_NOMAL | 0x38
	KC_CAPS_LOCK    = TYPE_NOMAL | 0x39
	KC_CLCK         = KC_CAPS_LOCK
	KC_F1           = TYPE_NOMAL | 0x3A
	KC_F2           = TYPE_NOMAL | 0x3B
	KC_F3           = TYPE_NOMAL | 0x3C
	KC_F4           = TYPE_NOMAL | 0x3D
	KC_F5           = TYPE_NOMAL | 0x3E
	KC_F6           = TYPE_NOMAL | 0x3F
	KC_F7           = TYPE_NOMAL | 0x40
	KC_F8           = TYPE_NOMAL | 0x41
	KC_F9           = TYPE_NOMAL | 0x42
	KC_F10          = TYPE_NOMAL | 0x43
	KC_F11          = TYPE_NOMAL | 0x44
	KC_F12          = TYPE_NOMAL | 0x45
	KC_PRINTSCREEN  = TYPE_NOMAL | 0x46
	KC_PS           = KC_PRINTSCREEN
	KC_SCROLL_LOCK  = TYPE_NOMAL | 0x47
	KC_SLCK         = KC_SCROLL_LOCK
	KC_PAUSE        = TYPE_NOMAL | 0x48
	KC_INSERT       = TYPE_NOMAL | 0x49
	KC_HOME         = TYPE_NOMAL | 0x4A
	KC_PAGE_UP      = TYPE_NOMAL | 0x4B
	KC_DELETE       = TYPE_NOMAL | 0x4C
	KC_END          = TYPE_NOMAL | 0x4D
	KC_PAGE_DOWN    = TYPE_NOMAL | 0x4E
	KC_RIGHT        = TYPE_NOMAL | 0x4F
	KC_LEFT         = TYPE_NOMAL | 0x50
	KC_DOWN         = TYPE_NOMAL | 0x51
	KC_UP           = TYPE_NOMAL | 0x52
	KC_NUM_LOCK     = TYPE_NOMAL | 0x53
	KC_NLCK         = KC_NUM_LOCK
	KC_PAD_SLASH    = TYPE_NOMAL | 0x54
	KC_PAD_ASTERISK = TYPE_NOMAL | 0x55
	KC_PAD_MINUS    = TYPE_NOMAL | 0x56
	KC_PAD_PLUS     = TYPE_NOMAL | 0x57
	KC_PAD_ENTER    = TYPE_NOMAL | 0x58
	KC_PAD1         = TYPE_NOMAL | 0x59
	KC_PAD2         = TYPE_NOMAL | 0x5A
	KC_PAD3         = TYPE_NOMAL | 0x5B
	KC_PAD4         = TYPE_NOMAL | 0x5C
	KC_PAD5         = TYPE_NOMAL | 0x5D
	KC_PAD6         = TYPE_NOMAL | 0x5E
	KC_PAD7         = TYPE_NOMAL | 0x5F
	KC_PAD8         = TYPE_NOMAL | 0x60
	KC_PAD9         = TYPE_NOMAL | 0x61
	KC_PAD0         = TYPE_NOMAL | 0x62
	KC_PAD_PERIOD   = TYPE_NOMAL | 0x63
	KC_NON_USBS     = TYPE_NOMAL | 0x64
	KC_MENU         = TYPE_NOMAL | 0x65
	KC_F13          = TYPE_NOMAL | 0x68
	KC_F14          = TYPE_NOMAL | 0x69
	KC_F15          = TYPE_NOMAL | 0x6A
	KC_F16          = TYPE_NOMAL | 0x6B
	KC_F17          = TYPE_NOMAL | 0x6C
	KC_F18          = TYPE_NOMAL | 0x6D
	KC_F19          = TYPE_NOMAL | 0x6E
	KC_F20          = TYPE_NOMAL | 0x6F
	KC_F21          = TYPE_NOMAL | 0x70
	KC_F22          = TYPE_NOMAL | 0x71
	KC_F23          = TYPE_NOMAL | 0x72
	KC_F24          = TYPE_NOMAL | 0x73
	KC_BACKSLASH    = TYPE_NOMAL | 0x87 // \ |
	KC_HIRAGANA     = TYPE_NOMAL | 0x88
	KC_BACKSLASH2   = TYPE_NOMAL | 0x89 // \ _
	KC_HENKAN       = TYPE_NOMAL | 0x8A
	KC_MUHENKAN     = TYPE_NOMAL | 0x8B
	KC_KANA         = TYPE_NOMAL | 0x90
	KC_EISU         = TYPE_NOMAL | 0x91
	KC_LEFT_CTRL    = TYPE_NOMAL | 0xE0
	KC_LCTRL        = KC_LEFT_CTRL
	KC_LEFT_SHIFT   = TYPE_NOMAL | 0xE1
	KC_LSHIFT       = KC_LEFT_SHIFT
	KC_LEFT_ALT     = TYPE_NOMAL | 0xE2
	KC_LALT         = KC_LEFT_ALT
	KC_WINDOWS      = TYPE_NOMAL | 0xE3
	KC_GUI          = KC_WINDOWS
	KC_RIGHT_CTRL   = TYPE_NOMAL | 0xE4
	KC_RCTRL        = KC_RIGHT_CTRL
	KC_RIGHT_SHIFT  = TYPE_NOMAL | 0xE5
	KC_RSHIFT       = KC_RIGHT_SHIFT
)

const (
	KC_MEDIA_PLAY         = TYPE_MEDIA | 0xB0
	KC_MEDIA_PAUSE        = TYPE_MEDIA | 0xB1
	KC_MEDIA_RECORD       = TYPE_MEDIA | 0xB2
	KC_MEDIA_FAST_FORWARD = TYPE_MEDIA | 0xB3
	KC_MEDIA_REWIND       = TYPE_MEDIA | 0xB4
	KC_MEDIA_NEXT_TRACK   = TYPE_MEDIA | 0xB5
	KC_MEDIA_PREV_TRACK   = TYPE_MEDIA | 0xB6
	KC_MEDIA_STOP         = TYPE_MEDIA | 0xB7
	KC_MEDIA_EJECT        = TYPE_MEDIA | 0xB8
	KC_MEDIA_RANDOM_PLAY  = TYPE_MEDIA | 0xB9
	KC_MEDIA_PLAY_PAUSE   = TYPE_MEDIA | 0xCD
	KC_MEDIA_PLAY_SKIP    = TYPE_MEDIA | 0xCE
	KC_MEDIA_MUTE         = TYPE_MEDIA | 0xE2
	KC_MEDIA_VOLUME_INC   = TYPE_MEDIA | 0xE9
	KC_MEDIA_VOLUME_DEC   = TYPE_MEDIA | 0xEA
)

const (
	MOUSE_LEFT       = TYPE_MOUSE | 0x01 // mouse.Left
	MOUSE_RIGHT      = TYPE_MOUSE | 0x02 // mouse.Right
	MOUSE_MIDDLE     = TYPE_MOUSE | 0x04 // mouse.Middle
	MOUSE_BACK       = TYPE_MOUSE | 0x08 // mouse.Back
	MOUSE_FORWARD    = TYPE_MOUSE | 0x10 // mouse.Forward
	MOUSE_WHEEL_DOWN = TYPE_MOUSE | 0x20
	MOUSE_WHEEL_UP   = TYPE_MOUSE | 0x40
)
