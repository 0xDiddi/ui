// 12 august 2018

package ui

// #include "ui.h"
import "C"

// Attribute stores information about an attribute in an
// AttributedString.
//
// The following types can be used as Attributes:
//
// 	- TextFamily
// 	- TextSize
// 	- TextWeight
// 	- TextItalic
// 	- TextStretch
// 	- TextColor
// 	- TextBackgroundColor
// 	- TextUnderline
// 	- TextUnderlineColor
// 	- OpenTypeFeatures
//
// For every Unicode codepoint in the AttributedString, at most one
// value of each attribute type can be applied.
type Attribute interface {
	toLibui() *C.uiAttribute
}

// TextFamily is an Attribute that changes the font family of the text
// it is applied to. Font family names are case-insensitive.
type TextFamily string

func (f TextFamily) toLibui() *C.uiAttribute {
	fstr := C.CString(string(f))
	defer freestr(fstr)
	return C.uiNewFamilyAttribute(fstr)
}

// TextSize is an Attribute that changes the size of the text it is
// applied to, in typographical points.
type TextSize float64

func (s TextSize) toLibui() *C.uiAttribute {
	return C.uiNewSizeAttribute(C.double(size))
}

// TextWeight is an Attribute that changes the weight of the text
// it is applied to. These roughly map to the OS/2 text weight field
// of TrueType and OpenType fonts, or to CSS weight numbers. The
// named constants are nominal values; the actual values may vary
// by font and by OS, though this isn't particularly likely. Any value
// between TextWeightMinimum and TextWeightMaximum,
// inclusive, is allowed.
//
// Note that due to restrictions in early versions of Windows, some
// fonts have "special" weights be exposed in many programs as
// separate font families. This is perhaps most notable with
// Arial Black. Package ui does not do this, even on Windows
// (because the DirectWrite API libui uses on Windows does not do
// this); to specify Arial Black, use family Arial and weight
// TextWeightBlack.
type TextWeight int
const (
	TextWeightMinimum = 0,
	TextWeightThin = 100,
	TextWeightUltraLight = 200,
	TextWeightLight = 300,
	TextWeightBook = 350,
	TextWeightNormal = 400,
	TextWeightMedium = 500,
	TextWeightSemiBold = 600,
	TextWeightBold = 700,
	TextWeightUltraBold = 800,
	TextWeightHeavy = 900,
	TextWeightUltraHeavy = 950,
	TextWeightMaximum = 1000,
)

func (w TextWeight) toLibui() *C.uiAttribute {
	return C.uiNewWeightAttribute(C.uiTextWeight(w))
}

// TextItalic is an Attribute that changes the italic mode of the text
// it is applied to. Italic represents "true" italics where the slanted
// glyphs have custom shapes, whereas oblique represents italics
// that are merely slanted versions of the normal glyphs. Most fonts
// usually have one or the other.
type TextItalic int
const (
	TextItalicNormal TextItalic = iota
	TextItalicOblique
	TextItalicItalic
)

func (i TextItalic) toLibui() *C.uiAttribute {
	return C.uiNewItalicAttribute(C.uiTextItalic(i))
}

///// TODOTODO

// TextStretch is an Attribute that changes the stretch (also called
// "width") of the text it is applied to.
//
// Note that due to restrictions in early versions of Windows, some
// fonts have "special" stretches be exposed in many programs as
// separate font families. This is perhaps most notable with
// Arial Condensed. Package ui does not do this, even on Windows
// (because the DirectWrite API package ui uses on Windows does
// not do this); to specify Arial Condensed, use family Arial and
// stretch TextStretchCondensed.
type TextStretch int
const (
	TextStretchUltraCondensed TextStretch = iota
	TextStretchExtraCondensed
	TextStretchCondensed
	TextStretchSemiCondensed
	TextStretchNormal
	TextStretchSemiExpanded
	TextStretchExpanded
	TextStretchExtraExpanded
	TextStretchUltraExpanded
)

func (s TextStretch) toLibui() *C.uiAttribute {
	return C.uiNewStretchAttribute(C.uiTextStretch(s))
}

/////// TODOTODO

// uiNewColorAttribute() creates a new uiAttribute that changes the
// color of the text it is applied to. It is an error to specify an invalid
// color.
_UI_EXTERN uiAttribute *uiNewColorAttribute(double r, double g, double b, double a);

// uiAttributeColor() returns the text color stored in a. It is an
// error to call this on a uiAttribute that does not hold a text color.
_UI_EXTERN void uiAttributeColor(const uiAttribute *a, double *r, double *g, double *b, double *alpha);

// uiNewBackgroundAttribute() creates a new uiAttribute that
// changes the background color of the text it is applied to. It is an
// error to specify an invalid color.
_UI_EXTERN uiAttribute *uiNewBackgroundAttribute(double r, double g, double b, double a);

// TODO reuse uiAttributeColor() for background colors, or make a new function...

// uiUnderline specifies a type of underline to use on text.
_UI_ENUM(uiUnderline) {
	uiUnderlineNone,
	uiUnderlineSingle,
	uiUnderlineDouble,
	uiUnderlineSuggestion,		// wavy or dotted underlines used for spelling/grammar checkers
};

// uiNewUnderlineAttribute() creates a new uiAttribute that changes
// the type of underline on the text it is applied to. It is an error to
// specify an underline type not specified in uiUnderline.
_UI_EXTERN uiAttribute *uiNewUnderlineAttribute(uiUnderline u);

// uiAttributeUnderline() returns the underline type stored in a. It is
// an error to call this on a uiAttribute that does not hold an underline
// style.
_UI_EXTERN uiUnderline uiAttributeUnderline(const uiAttribute *a);

// uiUnderlineColor specifies the color of any underline on the text it
// is applied to, regardless of the type of underline. In addition to
// being able to specify a custom color, you can explicitly specify
// platform-specific colors for suggestion underlines; to use them
// correctly, pair them with uiUnderlineSuggestion (though they can
// be used on other types of underline as well).
// 
// If an underline type is applied but no underline color is
// specified, the text color is used instead. If an underline color
// is specified without an underline type, the underline color
// attribute is ignored, but not removed from the uiAttributedString.
_UI_ENUM(uiUnderlineColor) {
	uiUnderlineColorCustom,
	uiUnderlineColorSpelling,
	uiUnderlineColorGrammar,
	uiUnderlineColorAuxiliary,		// for instance, the color used by smart replacements on macOS or in Microsoft Office
};

// uiNewUnderlineColorAttribute() creates a new uiAttribute that
// changes the color of the underline on the text it is applied to.
// It is an error to specify an underline color not specified in
// uiUnderlineColor.
//
// If the specified color type is uiUnderlineColorCustom, it is an
// error to specify an invalid color value. Otherwise, the color values
// are ignored and should be specified as zero.
_UI_EXTERN uiAttribute *uiNewUnderlineColorAttribute(uiUnderlineColor u, double r, double g, double b, double a);

// uiAttributeUnderlineColor() returns the underline color stored in
// a. It is an error to call this on a uiAttribute that does not hold an
// underline color.
_UI_EXTERN void uiAttributeUnderlineColor(const uiAttribute *a, uiUnderlineColor *u, double *r, double *g, double *b, double *alpha);

// uiOpenTypeFeatures represents a set of OpenType feature
// tag-value pairs, for applying OpenType features to text.
// OpenType feature tags are four-character codes defined by
// OpenType that cover things from design features like small
// caps and swashes to language-specific glyph shapes and
// beyond. Each tag may only appear once in any given
// uiOpenTypeFeatures instance. Each value is a 32-bit integer,
// often used as a Boolean flag, but sometimes as an index to choose
// a glyph shape to use.
// 
// If a font does not support a certain feature, that feature will be
// ignored. (TODO verify this on all OSs)
// 
// See the OpenType specification at
// https://www.microsoft.com/typography/otspec/featuretags.htm
// for the complete list of available features, information on specific
// features, and how to use them.
// TODO invalid features
typedef struct uiOpenTypeFeatures uiOpenTypeFeatures;

// uiOpenTypeFeaturesForEachFunc is the type of the function
// invoked by uiOpenTypeFeaturesForEach() for every OpenType
// feature in otf. Refer to that function's documentation for more
// details.
typedef uiForEach (*uiOpenTypeFeaturesForEachFunc)(const uiOpenTypeFeatures *otf, char a, char b, char c, char d, uint32_t value, void *data);

// @role uiOpenTypeFeatures constructor
// uiNewOpenTypeFeatures() returns a new uiOpenTypeFeatures
// instance, with no tags yet added.
_UI_EXTERN uiOpenTypeFeatures *uiNewOpenTypeFeatures(void);

// @role uiOpenTypeFeatures destructor
// uiFreeOpenTypeFeatures() frees otf.
_UI_EXTERN void uiFreeOpenTypeFeatures(uiOpenTypeFeatures *otf);

// uiOpenTypeFeaturesClone() makes a copy of otf and returns it.
// Changing one will not affect the other.
_UI_EXTERN uiOpenTypeFeatures *uiOpenTypeFeaturesClone(const uiOpenTypeFeatures *otf);

// uiOpenTypeFeaturesAdd() adds the given feature tag and value
// to otf. The feature tag is specified by a, b, c, and d. If there is
// already a value associated with the specified tag in otf, the old
// value is removed.
_UI_EXTERN void uiOpenTypeFeaturesAdd(uiOpenTypeFeatures *otf, char a, char b, char c, char d, uint32_t value);

// uiOpenTypeFeaturesRemove() removes the given feature tag
// and value from otf. If the tag is not present in otf,
// uiOpenTypeFeaturesRemove() does nothing.
_UI_EXTERN void uiOpenTypeFeaturesRemove(uiOpenTypeFeatures *otf, char a, char b, char c, char d);

// uiOpenTypeFeaturesGet() determines whether the given feature
// tag is present in otf. If it is, *value is set to the tag's value and
// nonzero is returned. Otherwise, zero is returned.
// 
// Note that if uiOpenTypeFeaturesGet() returns zero, value isn't
// changed. This is important: if a feature is not present in a
// uiOpenTypeFeatures, the feature is NOT treated as if its
// value was zero anyway. Script-specific font shaping rules and
// font-specific feature settings may use a different default value
// for a feature. You should likewise not treat a missing feature as
// having a value of zero either. Instead, a missing feature should
// be treated as having some unspecified default value.
_UI_EXTERN int uiOpenTypeFeaturesGet(const uiOpenTypeFeatures *otf, char a, char b, char c, char d, uint32_t *value);

// uiOpenTypeFeaturesForEach() executes f for every tag-value
// pair in otf. The enumeration order is unspecified. You cannot
// modify otf while uiOpenTypeFeaturesForEach() is running.
_UI_EXTERN void uiOpenTypeFeaturesForEach(const uiOpenTypeFeatures *otf, uiOpenTypeFeaturesForEachFunc f, void *data);

// uiNewFeaturesAttribute() creates a new uiAttribute that changes
// the font family of the text it is applied to. otf is copied; you may
// free it after uiNewFeaturesAttribute() returns.
_UI_EXTERN uiAttribute *uiNewFeaturesAttribute(const uiOpenTypeFeatures *otf);

// uiAttributeFeatures() returns the OpenType features stored in a.
// The returned uiOpenTypeFeatures object is owned by a. It is an
// error to call this on a uiAttribute that does not hold OpenType
// features.
_UI_EXTERN const uiOpenTypeFeatures *uiAttributeFeatures(const uiAttribute *a);

// uiAttributedString represents a string of UTF-8 text that can
// optionally be embellished with formatting attributes. libui
// provides the list of formatting attributes, which cover common
// formatting traits like boldface and color as well as advanced
// typographical features provided by OpenType like superscripts
// and small caps. These attributes can be combined in a variety of
// ways.
//
// Attributes are applied to runs of Unicode codepoints in the string.
// Zero-length runs are elided. Consecutive runs that have the same
// attribute type and value are merged. Each attribute is independent
// of each other attribute; overlapping attributes of different types
// do not split each other apart, but different values of the same
// attribute type do.
//
// The empty string can also be represented by uiAttributedString,
// but because of the no-zero-length-attribute rule, it will not have
// attributes.
//
// A uiAttributedString takes ownership of all attributes given to
// it, as it may need to duplicate or delete uiAttribute objects at
// any time. By extension, when you free a uiAttributedString,
// all uiAttributes within will also be freed. Each method will
// describe its own rules in more details.
//
// In addition, uiAttributedString provides facilities for moving
// between grapheme clusters, which represent a character
// from the point of view of the end user. The cursor of a text editor
// is always placed on a grapheme boundary, so you can use these
// features to move the cursor left or right by one "character".
// TODO does uiAttributedString itself need this
//
// uiAttributedString does not provide enough information to be able
// to draw itself onto a uiDrawContext or respond to user actions.
// In order to do that, you'll need to use a uiDrawTextLayout, which
// is built from the combination of a uiAttributedString and a set of
// layout-specific properties.
typedef struct uiAttributedString uiAttributedString;

// uiAttributedStringForEachAttributeFunc is the type of the function
// invoked by uiAttributedStringForEachAttribute() for every
// attribute in s. Refer to that function's documentation for more
// details.
typedef uiForEach (*uiAttributedStringForEachAttributeFunc)(const uiAttributedString *s, const uiAttribute *a, size_t start, size_t end, void *data);

// @role uiAttributedString constructor
// uiNewAttributedString() creates a new uiAttributedString from
// initialString. The string will be entirely unattributed.
_UI_EXTERN uiAttributedString *uiNewAttributedString(const char *initialString);

// @role uiAttributedString destructor
// uiFreeAttributedString() destroys the uiAttributedString s.
// It will also free all uiAttributes within.
_UI_EXTERN void uiFreeAttributedString(uiAttributedString *s);

// uiAttributedStringString() returns the textual content of s as a
// '\0'-terminated UTF-8 string. The returned pointer is valid until
// the next change to the textual content of s.
_UI_EXTERN const char *uiAttributedStringString(const uiAttributedString *s);

// uiAttributedStringLength() returns the number of UTF-8 bytes in
// the textual content of s, excluding the terminating '\0'.
_UI_EXTERN size_t uiAttributedStringLen(const uiAttributedString *s);

// uiAttributedStringAppendUnattributed() adds the '\0'-terminated
// UTF-8 string str to the end of s. The new substring will be
// unattributed.
_UI_EXTERN void uiAttributedStringAppendUnattributed(uiAttributedString *s, const char *str);

// uiAttributedStringInsertAtUnattributed() adds the '\0'-terminated
// UTF-8 string str to s at the byte position specified by at. The new
// substring will be unattributed; existing attributes will be moved
// along with their text.
_UI_EXTERN void uiAttributedStringInsertAtUnattributed(uiAttributedString *s, const char *str, size_t at);

// TODO add the Append and InsertAtExtendingAttributes functions
// TODO and add functions that take a string + length

// uiAttributedStringDelete() deletes the characters and attributes of
// s in the byte range [start, end).
_UI_EXTERN void uiAttributedStringDelete(uiAttributedString *s, size_t start, size_t end);

// TODO add a function to uiAttributedString to get an attribute's value at a specific index or in a specific range, so we can edit

// uiAttributedStringSetAttribute() sets a in the byte range [start, end)
// of s. Any existing attributes in that byte range of the same type are
// removed. s takes ownership of a; you should not use it after
// uiAttributedStringSetAttribute() returns.
_UI_EXTERN void uiAttributedStringSetAttribute(uiAttributedString *s, uiAttribute *a, size_t start, size_t end);

// uiAttributedStringForEachAttribute() enumerates all the
// uiAttributes in s. It is an error to modify s in f. Within f, s still
// owns the attribute; you can neither free it nor save it for later
// use.
// TODO reword the above for consistency (TODO and find out what I meant by that)
// TODO define an enumeration order (or mark it as undefined); also define how consecutive runs of identical attributes are handled here and sync with the definition of uiAttributedString itself
_UI_EXTERN void uiAttributedStringForEachAttribute(const uiAttributedString *s, uiAttributedStringForEachAttributeFunc f, void *data);

// TODO const correct this somehow (the implementation needs to mutate the structure)
_UI_EXTERN size_t uiAttributedStringNumGraphemes(uiAttributedString *s);

// TODO const correct this somehow (the implementation needs to mutate the structure)
_UI_EXTERN size_t uiAttributedStringByteIndexToGrapheme(uiAttributedString *s, size_t pos);

// TODO const correct this somehow (the implementation needs to mutate the structure)
_UI_EXTERN size_t uiAttributedStringGraphemeToByteIndex(uiAttributedString *s, size_t pos);

// uiFontDescriptor provides a complete description of a font where
// one is needed. Currently, this means as the default font of a
// uiDrawTextLayout and as the data returned by uiFontButton.
// All the members operate like the respective uiAttributes.
typedef struct uiFontDescriptor uiFontDescriptor;

struct uiFontDescriptor {
	// TODO const-correct this or figure out how to deal with this when getting a value
	char *Family;
	double Size;
	uiTextWeight Weight;
	uiTextItalic Italic;
	uiTextStretch Stretch;
};

// uiDrawTextLayout is a concrete representation of a
// uiAttributedString that can be displayed in a uiDrawContext.
// It includes information important for the drawing of a block of
// text, including the bounding box to wrap the text within, the
// alignment of lines of text within that box, areas to mark as
// being selected, and other things.
//
// Unlike uiAttributedString, the content of a uiDrawTextLayout is
// immutable once it has been created.
//
// TODO talk about OS-specific differences with text drawing that libui can't account for...
typedef struct uiDrawTextLayout uiDrawTextLayout;

// uiDrawTextAlign specifies the alignment of lines of text in a
// uiDrawTextLayout.
// TODO should this really have Draw in the name?
_UI_ENUM(uiDrawTextAlign) {
	uiDrawTextAlignLeft,
	uiDrawTextAlignCenter,
	uiDrawTextAlignRight,
};

// uiDrawTextLayoutParams describes a uiDrawTextLayout.
// DefaultFont is used to render any text that is not attributed
// sufficiently in String. Width determines the width of the bounding
// box of the text; the height is determined automatically.
typedef struct uiDrawTextLayoutParams uiDrawTextLayoutParams;

// TODO const-correct this somehow
struct uiDrawTextLayoutParams {
	uiAttributedString *String;
	uiFontDescriptor *DefaultFont;
	double Width;
	uiDrawTextAlign Align;
};

// @role uiDrawTextLayout constructor
// uiDrawNewTextLayout() creates a new uiDrawTextLayout from
// the given parameters.
//
// TODO
// - allow creating a layout out of a substring
// - allow marking compositon strings
// - allow marking selections, even after creation
// - add the following functions:
// 	- uiDrawTextLayoutHeightForWidth() (returns the height that a layout would need to be to display the entire string at a given width)
// 	- uiDrawTextLayoutRangeForSize() (returns what substring would fit in a given size)
// 	- uiDrawTextLayoutNewWithHeight() (limits amount of string used by the height)
// - some function to fix up a range (for text editing)
_UI_EXTERN uiDrawTextLayout *uiDrawNewTextLayout(uiDrawTextLayoutParams *params);

// @role uiDrawFreeTextLayout destructor
// uiDrawFreeTextLayout() frees tl. The underlying
// uiAttributedString is not freed.
_UI_EXTERN void uiDrawFreeTextLayout(uiDrawTextLayout *tl);

// uiDrawText() draws tl in c with the top-left point of tl at (x, y).
_UI_EXTERN void uiDrawText(uiDrawContext *c, uiDrawTextLayout *tl, double x, double y);

// uiDrawTextLayoutExtents() returns the width and height of tl
// in width and height. The returned width may be smaller than
// the width passed into uiDrawNewTextLayout() depending on
// how the text in tl is wrapped. Therefore, you can use this
// function to get the actual size of the text layout.
_UI_EXTERN void uiDrawTextLayoutExtents(uiDrawTextLayout *tl, double *width, double *height);
