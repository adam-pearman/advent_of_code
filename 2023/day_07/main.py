with open('day_07/input.txt') as f:
    lines = f.readlines()

hands = {
    'high': [],
    'pair': [],
    'two_pair': [],
    'three': [],
    'full_house': [],
    'four': [],
    'five': [],
}

# Part 1
# face_cards = {
#     'A': 14,
#     'K': 13,
#     'Q': 12,
#     'J': 11,
#     'T': 10,
# }

# def get_hand_type(hand):
#     cards = list(set(hand))
#     if len(cards) == 1:
#         return 'five'
#     if len(cards) == 2:
#         if hand.count(cards[0]) == 4 or hand.count(cards[0]) == 1:
#             return 'four'
#         elif hand.count(cards[0]) == 3  or hand.count(cards[0]) == 2:
#             return 'full_house'
#     if len(cards) == 3:
#         for card in cards:
#             if hand.count(card) == 3:
#                 return 'three'
#             elif hand.count(card) == 2:
#                 return 'two_pair'
#     if len(cards) == 4:
#         return 'pair'
#     if len(cards) == 5:
#         return 'high'

# Part 2
face_cards = {
    'A': 14,
    'K': 13,
    'Q': 12,
    'J': 1,
    'T': 10,
}

def get_hand_type(hand):
    cards = list(set(hand.replace('J', '')))
    j_count = hand.count('J')
    if len(cards) < 2:
        return 'five'
    if len(cards) == 2:
        if hand.count(cards[0]) == 4 - j_count or hand.count(cards[0]) == 1:
            return 'four'
        elif (hand.count(cards[0]) == 3 and j_count == 0) or (hand.count(cards[0]) == 2 and j_count < 2):
            return 'full_house'
    if len(cards) == 3:
        for card in cards:
            if hand.count(card) == 3 - j_count:
                return 'three'
            elif hand.count(card) == 2 and j_count < 1:
                return 'two_pair'
    if len(cards) == 4:
        return 'pair'
    if len(cards) == 5:
        return 'high'


def rank_hand(hand_type, hand, bid):
    if len(hands[hand_type]) == 0:
        hands[hand_type].append((hand, bid))
    else:
        for index, ranked_hand in enumerate(hands[hand_type]):
            for char in range(5):
                ranked_score = int(ranked_hand[0][char]) if ranked_hand[0][char] not in face_cards else face_cards[ranked_hand[0][char]]
                hand_score = int(hand[char]) if hand[char] not in face_cards else face_cards[hand[char]]
                if hand_score < ranked_score:
                    hands[hand_type].insert(index, (hand, bid))
                    return
                elif hand_score > ranked_score:
                    break
        hands[hand_type].append((hand, bid))


for line in lines:
    hand, bid = line.strip().split(' ')
    hand_type = get_hand_type(hand)
    rank_hand(hand_type, hand, int(bid))

ranked_hands = [hand for hand_type in hands.values() for hand in hand_type]

winnings = 0
for index, hand in enumerate(ranked_hands):
    winnings += (index + 1) * hand[1]

print(winnings)
