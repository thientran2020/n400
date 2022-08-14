package models

const (
	LINK            = "uscis.gov/citizenship/testupdates"
	PROMT           = "Enter to continue. [Q] to quit: "
	WELCOME_MESSAGE = "\n_________WELCOME TO PREPATION N-400 TEST_________\n\n"
	EXIT_MESSAGE    = "\nGOOD LUCK WITH YOUR EXAM ^.^\n"
	RED             = "\033[31m"
	BLUE            = "\033[34m"
	GREEN           = "\033[32m"
	YELLOW          = "\033[33m"
	UWHITE          = "\033[4;37m"
	COLOROFF        = "\033[0m"
)

var (
	CATEGORY_KEYS    = []int{1, 58, 88}
	SUBCATEGORY_KEYS = []int{1, 13, 48, 58, 71, 78, 88, 96, 99}
)

var CATEGORIES = map[int]string{
	1:  "American Government",
	58: "American History",
	88: "Integrated Civics",
}

var SUBCATEGORIES = map[int]string{
	1:  "Principles of American Democracy",
	13: "System of Government",
	48: "Right and Responsibilities",
	58: "Colonial Period and Independence",
	71: "1800's",
	78: "Recent American History and Other Important Historical Information",
	88: "Geography",
	96: "Symbols",
	99: "Holidays",
}

var QUESTIONS = []string{
	"What is the supreme law of the land?",
	"What does the Constitution do?",
	"The idea of self-government is in the first 3 words of the Constitution. What are these words?",
	"What is an amendment?",
	"What do we call the first 10 amendments to the Constitution?",
	"What is one right or freedom from the First Amendment?",
	"How many amendments does the Constitution have?",
	"What did the Declaration of Independence do?",
	"What are 2 rights in the Declaration of Independence?",
	"What os freedom of religion?",
	"What is the economic system in the U.S.?",
	"What is the 'rule of law'?",
	// -------------------------- ///
	"Name 1 branch or part of the government?",
	"What stops 1 branch of government from becoming too powerful?",
	"Who is in charge of the executive branch?",
	"Who makes federal laws?",
	"What are the 2 parts of the U.S. Congress?",
	"How many U.S. Senators are there?",
	"We elect a U.S. Senator for how many years?",
	"Who is 1 of your state's U.S. Senators now? Dianne Feinstein || Alex Padilla", // ***
	"The House of Representatives has how many voting members?",
	"We elect a U.S. Representative for how many years?",
	"Name your U.S. Representative. Ro Khanna", // ***
	"Who does a U.S. Senator present?",
	"Why do some states have more Representatives than other states?",
	"We elect a President for how many years?",
	"In what month do we vote for President?",
	"What is the name of the President of the U.S. now? Joe Biden",          // ***
	"What is the name of the Vice President of the U.S. now? Kamala Harris", // ***
	"If the President can no longer serve, who becomes President?",
	"If both the President and the Vice President can no longer serve, who becomes President?",
	"Who is the Commander in Chief of the military?",
	"Who signs bills to become laws?",
	"Who vetoes bills?",
	"What does the President's Cabinet do?",
	"What are 2 Cabinet-level positions?",
	"What does the judicial branch do?",
	"What is the highest court in the U.S.?",
	"How many justices are on the Supreme Court?",
	"Who is the Chief Justice of the U.S. now? John Roberts", // ***
	"Under our Constitution, some powers belong to the federal government. What is 1 power of the federal government?",
	"Under our Constitution, some powers belong to the states. What is 1 power of the states?",
	"Who is the Governor of your state now? Gavin Newsom", // ***
	"What is the capital of your state? Sacramento",       // ***
	"What are the 2 major political parties in the U.S.?",
	"What is the political party of the President now? Democratic Party",                // ***
	"What is the name of the Speaker of the House of Representatives now? Nancy Pelosi", // ***
	// -------------------------- ///
	"There are 4 amendments to the Constitution about who can vote. Describe 1 of them.",
	"What is 1 responsibility that is only for U.S. citizens?",
	"Name 1 right only for U.S. citizens.",
	"What are 2 rights of everyone living in the U.S.?",
	"What do we show loyalty to when we say the Pledge of Allegiance?",
	"What is one promise you make when you become a U.S. citizen?",
	"How old do citizens have to be to vote for President?",
	"What are 2 ways that Americans can participate in their democracy?",
	"When is the last day you can send in federal income tax forms?",
	"When must all men register for the Selective Service?",
	// -------------------------- ///
	"What is 1 reason colonists came to America?",
	"Who lived in America before the Europeans arrived?",
	"What group of people was taken to America and sold as slaves?",
	"Why did the colonists fight the British?",
	"Who wrote the Declaration of Independence?",
	"When was the Declaration of Independence adopted?",
	"There were 13 original states. Name 3?",
	"What happened at the Consitutional Convention?",
	"When was the Constitution written?",
	"The Federalist Papers supported the passage of the U.S. Constitution. Name 1 of the writers.",
	"What is 1 thing Benjamin Franklin is famous for?",
	"Who is the 'Father of Our Country'?",
	"Who was the first President?",
	// -------------------------- ///
	"What territory did the U.S. buy from France in 1803?",
	"Name 1 war fought by the U.S. in the 1800s.",
	"Name the U.S. war between the North and the South.",
	"Name 1 problem that led to the Civil War.",
	"What was 1 important thing that Abraham Lincoln did?",
	"What did the Emancipation Proclamation do?",
	"What did Susan B. Anthony do?",
	// -------------------------- ///
	"Name 1 war fought by the U.S. in the 1900s.",
	"Who was President during World War I?",
	"Who was President during the Great Depression and World War II?",
	"Who did the U.S. fight in World War II?",
	"Before he was President, Eisenhower was a general. What war he was in?",
	"During the Cold War, what was the main concern of the U.S.?",
	"What movement tried to end racial discrimination?",
	"What did Martin Luther King, Jr. do?",
	"What major event happend on September 11, 2001, in the U.S.?",
	"Name 1 American Indian tribe in the U.S.",
	// -------------------------- ///
	"Name 1 of the 2 longest rivers in the U.S.",
	"What ocean is on the West Coast of the U.S.?",
	"What ocean is on the East Coast of the U.S.?",
	"Name 1 U.S. territory",
	"Name 1 state that borders Canada.",
	"Name 1 state that borders Mexico.",
	"What is the capital of the U.S.?",
	"Where is the Statue of Liberty?",
	// -------------------------- ///
	"Why does the flag have 13 stripes?",
	"Why does the flag have 50 stars?",
	"What is the name of the national anthem?",
	// -------------------------- ///
	"When do we celebrate Independence Day?",
	"Name 2 national U.S. holidays?",
}
