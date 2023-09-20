package geospacial

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountDistante(t *testing.T) {
	point := [][]float64{{1.348154, 103.680842}, {1.348109, 103.680764}, {1.348063, 103.680687}, {1.348018, 103.680609}, {1.347974, 103.68053}, {1.34793, 103.680452}, {1.347886, 103.680373}, {1.347841, 103.680295}, {1.347797, 103.680216}, {1.347752, 103.680138}, {1.347707, 103.68006}, {1.347662, 103.679982}, {1.347622, 103.679901}, {1.347583, 103.67982}, {1.347545, 103.679738}, {1.347507, 103.679656}, {1.34748, 103.679573}, {1.347433, 103.679496}, {1.347386, 103.679419}, {1.347344, 103.679339}, {1.347305, 103.679258}, {1.347269, 103.679175}, {1.347236, 103.679091}, {1.347206, 103.679007}, {1.347191, 103.678918}, {1.34718, 103.678829}, {1.34717, 103.678739}, {1.347157, 103.67865}, {1.34715, 103.67856}, {1.347132, 103.678472}, {1.34712, 103.678382}, {1.347033, 103.678391}, {1.346943, 103.678395}, {1.346853, 103.678398}, {1.346764, 103.678408}, {1.346675, 103.678423}, {1.346586, 103.678437}, {1.346497, 103.678453}, {1.346408, 103.678467}, {1.346318, 103.678475}, {1.346228, 103.67848}, {1.346139, 103.678486}, {1.346049, 103.678489}, {1.345958, 103.678486}, {1.345868, 103.678482}, {1.345778, 103.678479}, {1.345688, 103.678472}, {1.345599, 103.678466}, {1.345509, 103.67846}, {1.345419, 103.678454}, {1.345329, 103.67845}, {1.345239, 103.67845}, {1.345148, 103.67845}, {1.345059, 103.67846}, {1.344973, 103.678489}, {1.344893, 103.678529}, {1.344809, 103.678562}, {1.344726, 103.678596}, {1.344641, 103.678627}, {1.344557, 103.678658}, {1.344473, 103.678692}, {1.34439, 103.678727}, {1.344307, 103.678761}, {1.344224, 103.678796}, {1.34414, 103.67883}, {1.344056, 103.678863}, {1.343972, 103.678896}, {1.343887, 103.678925}, {1.343799, 103.678944}, {1.343711, 103.678963}, {1.343622, 103.678977}, {1.343533, 103.678988}, {1.343443, 103.678999}, {1.343354, 103.67901}, {1.343264, 103.679021}, {1.343175, 103.679032}, {1.343086, 103.679043}, {1.342996, 103.679054}, {1.342906, 103.679063}, {1.342817, 103.679073}, {1.342727, 103.679083}, {1.342638, 103.679092}, {1.342548, 103.679102}, {1.342458, 103.679111}, {1.342369, 103.679121}, {1.342279, 103.679131}, {1.34219, 103.67914}, {1.3421, 103.67915}, {1.34201, 103.679145}, {1.34192, 103.67914}, {1.34183, 103.679135}, {1.34174, 103.67913}, {1.34165, 103.679125}, {1.34156, 103.67912}, {1.34147, 103.67912}, {1.34138, 103.67912}, {1.34129, 103.67912}, {1.341199, 103.67912}, {1.341109, 103.67912}, {1.341019, 103.67912}, {1.340929, 103.67912}, {1.340841, 103.679135}, {1.340758, 103.679169}, {1.340683, 103.679217}, {1.340619, 103.679281}, {1.340555, 103.679345}, {1.340492, 103.679408}, {1.340428, 103.679472}, {1.340364, 103.679536}, {1.3403, 103.6796}, {1.340237, 103.679663}, {1.340173, 103.679727}, {1.340109, 103.679791}, {1.340045, 103.679855}, {1.339985, 103.679921}, {1.339951, 103.680004}, {1.339916, 103.680087}, {1.339889, 103.680172}, {1.339885, 103.680262}, {1.339881, 103.680352}, {1.339877, 103.680442}, {1.339873, 103.680532}, {1.339869, 103.680622}, {1.339865, 103.680712}, {1.339861, 103.680802}, {1.339867, 103.680892}, {1.339877, 103.680982}, {1.339886, 103.681071}, {1.339896, 103.681161}, {1.339906, 103.68125}, {1.339916, 103.68134}, {1.339926, 103.68143}, {1.339936, 103.681519}, {1.33995, 103.681608}, {1.33997, 103.681696}, {1.339989, 103.681784}, {1.340008, 103.681872}, {1.340027, 103.68196}, {1.340047, 103.682048}, {1.34007, 103.682135}, {1.340103, 103.682219}, {1.34014, 103.682301}, {1.340178, 103.682383}, {1.340217, 103.682464}, {1.340256, 103.682545}, {1.340295, 103.682627}, {1.340336, 103.682707}, {1.340382, 103.682784}, {1.340428, 103.682862}, {1.340474, 103.682939}, {1.340525, 103.683013}, {1.340578, 103.683086}, {1.340631, 103.683159}, {1.340684, 103.683232}, {1.34074, 103.683303}, {1.340795, 103.683374}, {1.34085, 103.683445}, {1.340906, 103.683516}, {1.340965, 103.683584}, {1.341021, 103.683654}, {1.341078, 103.683724}, {1.34114, 103.68379}, {1.341201, 103.683856}, {1.341262, 103.683922}, {1.341324, 103.683988}, {1.341385, 103.684054}, {1.341442, 103.684124}, {1.341497, 103.684195}, {1.341553, 103.684266}, {1.341608, 103.684338}, {1.341663, 103.684409}, {1.341717, 103.684481}, {1.341773, 103.684551}, {1.34183, 103.684622}, {1.341886, 103.684692}, {1.341943, 103.684762}, {1.341999, 103.684832}, {1.342056, 103.684903}, {1.342112, 103.684973}, {1.342169, 103.685043}, {1.342225, 103.685113}, {1.342282, 103.685183}, {1.342338, 103.685254}, {1.342395, 103.685324}, {1.342451, 103.685394}, {1.342508, 103.685464}, {1.342564, 103.685535}, {1.342621, 103.685605}, {1.342677, 103.685675}, {1.342734, 103.685745}, {1.34279, 103.685815}, {1.342847, 103.685886}, {1.342903, 103.685956}, {1.34296, 103.686026}, {1.343016, 103.686096}, {1.343073, 103.686167}, {1.343129, 103.686237}, {1.343186, 103.686307}, {1.343244, 103.686376}, {1.343302, 103.686444}, {1.34336, 103.686513}, {1.343418, 103.686583}, {1.343472, 103.686655}, {1.343527, 103.686726}, {1.343582, 103.686798}, {1.343636, 103.68687}, {1.343691, 103.686941}, {1.343745, 103.687013}, {1.3438, 103.687085}, {1.343854, 103.687157}, {1.343894, 103.687238}, {1.343942, 103.687313}, {1.34399, 103.687389}, {1.34404, 103.687464}, {1.344094, 103.687536}, {1.344147, 103.687609}, {1.344201, 103.687681}, {1.34426, 103.687749}, {1.34432, 103.687817}, {1.344379, 103.687885}, {1.344438, 103.687952}, {1.344498, 103.68802}, {1.344556, 103.688089}, {1.344615, 103.688157}, {1.344674, 103.688226}, {1.344739, 103.688281}, {1.344829, 103.688286}, {1.344919, 103.688288}, {1.345008, 103.688281}, {1.34509, 103.688244}, {1.34517, 103.688202}, {1.34525, 103.688161}, {1.345327, 103.688114}, {1.345402, 103.688065}, {1.345477, 103.688014}, {1.345551, 103.687963}, {1.345622, 103.687908}, {1.345691, 103.68785}, {1.34576, 103.687792}, {1.345828, 103.687733}, {1.345898, 103.687675}, {1.345967, 103.687618}, {1.346036, 103.68756}, {1.346105, 103.687502}, {1.346175, 103.687445}, {1.346235, 103.687378}, {1.346295, 103.687311}, {1.346355, 103.687244}, {1.346415, 103.687176}, {1.346476, 103.687109}, {1.346536, 103.687042}, {1.346596, 103.686975}, {1.346654, 103.686906}, {1.346713, 103.686838}, {1.346771, 103.686769}, {1.346838, 103.686709}, {1.34691, 103.686655}, {1.346986, 103.686608}, {1.347072, 103.686579}, {1.34716, 103.686563}, {1.34725, 103.686557}, {1.34734, 103.686552}, {1.34743, 103.68655}, {1.34752, 103.68655}, {1.34761, 103.686555}, {1.3477, 103.686559}, {1.34779, 103.686564}, {1.34788, 103.686568}, {1.34797, 103.686576}, {1.34806, 103.686585}, {1.348149, 103.686585}, {1.348239, 103.686574}, {1.348328, 103.686567}, {1.348411, 103.686532}, {1.348494, 103.686497}, {1.348569, 103.686447}, {1.348643, 103.686396}, {1.34869, 103.686319}, {1.348738, 103.686243}, {1.348764, 103.686157}, {1.34879, 103.68607}, {1.348814, 103.685984}, {1.348838, 103.685897}, {1.348824, 103.685808}, {1.348825, 103.685719}, {1.348847, 103.685632}, {1.348877, 103.685549}, {1.348951, 103.685507}, {1.349023, 103.685462}, {1.349085, 103.685397}, {1.349151, 103.685335}, {1.349218, 103.685275}, {1.349291, 103.685223}, {1.349348, 103.685155}, {1.349423, 103.685106}, {1.349496, 103.685053}, {1.349518, 103.684969}, {1.349592, 103.68493}, {1.349674, 103.684956}, {1.3497, 103.685039}, {1.349719, 103.685118}, {1.349772, 103.685191}, {1.349824, 103.685264}, {1.349874, 103.685339}, {1.349938, 103.685402}, {1.350007, 103.68546}, {1.35008, 103.685513}, {1.350153, 103.685565}, {1.350231, 103.685611}, {1.350313, 103.685647}, {1.350397, 103.68568}, {1.35048, 103.685714}, {1.350564, 103.685748}, {1.350653, 103.685759}, {1.350743, 103.685766}, {1.350833, 103.685772}, {1.350923, 103.685777}, {1.351013, 103.68578}, {1.351103, 103.68578}, {1.351193, 103.68578}, {1.351284, 103.68578}, {1.351374, 103.68578}, {1.351464, 103.68578}, {1.351554, 103.68578}, {1.351644, 103.68578}, {1.351734, 103.68578}, {1.351824, 103.685782}, {1.351914, 103.685786}, {1.352004, 103.685793}, {1.352092, 103.685814}, {1.352178, 103.68584}, {1.352264, 103.685867}, {1.35235, 103.685894}, {1.352432, 103.685931}, {1.352513, 103.685971}, {1.352593, 103.686012}, {1.352673, 103.686053}, {1.352753, 103.686094}, {1.352833, 103.686136}, {1.352913, 103.686178}, {1.352993, 103.68622}, {1.353072, 103.686262}, {1.353152, 103.686304}, {1.353232, 103.686347}, {1.353311, 103.686389}, {1.353391, 103.686431}, {1.35347, 103.686474}, {1.353549, 103.686517}, {1.353626, 103.686564}, {1.353704, 103.68661}, {1.35378, 103.686658}, {1.353856, 103.686706}, {1.353933, 103.686754}, {1.354008, 103.686803}, {1.354083, 103.686853}, {1.354158, 103.686904}, {1.354232, 103.686954}, {1.354306, 103.687007}, {1.354378, 103.68706}, {1.354451, 103.687113}, {1.354524, 103.687166}, {1.354596, 103.68722}, {1.354667, 103.687276}, {1.354733, 103.687336}, {1.354792, 103.687404}, {1.354852, 103.687472}, {1.354909, 103.687542}, {1.354966, 103.687611}, {1.355024, 103.687681}, {1.355081, 103.68775}, {1.355139, 103.68782}, {1.355202, 103.687859}, {1.355289, 103.68786}, {1.355358, 103.687842}, {1.355417, 103.687774}, {1.355476, 103.687706}, {1.355526, 103.687632}, {1.355582, 103.687561}, {1.355636, 103.687489}, {1.355683, 103.687412}, {1.355734, 103.687338}, {1.355776, 103.687258}, {1.355816, 103.687178}, {1.355856, 103.687097}, {1.355889, 103.687013}, {1.355916, 103.686927}, {1.35594, 103.68684}, {1.355951, 103.686751}, {1.35596, 103.686661}, {1.35596, 103.686571}, {1.35596, 103.686481}, {1.355955, 103.686391}, {1.355948, 103.686301}, {1.35594, 103.686212}, {1.355928, 103.686122}, {1.355918, 103.686033}, {1.355908, 103.685943}, {1.355888, 103.685855}, {1.355863, 103.685769}, {1.355832, 103.685684}, {1.355792, 103.685604}, {1.355746, 103.685526}, {1.355698, 103.68545}, {1.355652, 103.685373}, {1.3556, 103.685299}, {1.355549, 103.685225}, {1.355497, 103.685151}, {1.355452, 103.685073}, {1.355407, 103.684995}, {1.355362, 103.684917}, {1.355314, 103.68484}, {1.355266, 103.684764}, {1.355215, 103.68469}, {1.355166, 103.684614}, {1.35512, 103.684537}, {1.355071, 103.684462}, {1.355021, 103.684386}, {1.354972, 103.684311}, {1.354923, 103.684235}, {1.354876, 103.684158}, {1.354828, 103.684082}, {1.354777, 103.684008}, {1.354725, 103.683934}, {1.354674, 103.68386}, {1.354623, 103.683785}, {1.354573, 103.68371}, {1.354523, 103.683635}, {1.354473, 103.68356}, {1.354423, 103.683485}, {1.354373, 103.68341}, {1.354322, 103.683336}, {1.35427, 103.683263}, {1.354217, 103.683189}, {1.354165, 103.683116}, {1.354113, 103.683042}, {1.354059, 103.68297}, {1.354004, 103.682899}, {1.35395, 103.682827}, {1.353895, 103.682756}, {1.35384, 103.682684}, {1.353785, 103.682613}, {1.35373, 103.682541}, {1.353675, 103.68247}, {1.353619, 103.682399}, {1.353563, 103.682328}, {1.353507, 103.682258}, {1.353451, 103.682187}, {1.353397, 103.682115}, {1.353343, 103.682043}, {1.353289, 103.681971}, {1.353236, 103.681898}, {1.353182, 103.681826}, {1.353127, 103.681754}, {1.353072, 103.681683}, {1.353018, 103.681611}, {1.352963, 103.68154}, {1.352905, 103.68147}, {1.352846, 103.681402}, {1.352787, 103.681335}, {1.352727, 103.681267}, {1.352661, 103.681206}, {1.352595, 103.681145}, {1.352529, 103.681083}, {1.352464, 103.681021}, {1.352398, 103.680959}, {1.35233, 103.6809}, {1.352262, 103.680841}, {1.352194, 103.680781}, {1.352127, 103.680722}, {1.352059, 103.680663}, {1.351991, 103.680603}, {1.351923, 103.680544}, {1.351855, 103.680485}, {1.351787, 103.680426}, {1.351717, 103.680368}, {1.351652, 103.680307}, {1.351586, 103.680245}, {1.35152, 103.680184}, {1.351454, 103.680122}, {1.351387, 103.680062}, {1.351318, 103.680004}, {1.351249, 103.679946}, {1.351181, 103.679887}, {1.351112, 103.679829}, {1.351043, 103.679771}, {1.350989, 103.679821}, {1.350939, 103.679896}, {1.350889, 103.679971}, {1.350839, 103.680046}, {1.350789, 103.680121}, {1.350739, 103.680196}, {1.350689, 103.680271}, {1.350639, 103.680346}, {1.350589, 103.680421}, {1.350526, 103.680484}, {1.350462, 103.680548}, {1.350398, 103.680612}, {1.350334, 103.680676}, {1.350271, 103.680739}, {1.350202, 103.680797}, {1.350132, 103.680855}, {1.350063, 103.680912}, {1.349993, 103.68097}, {1.349923, 103.681027}, {1.349852, 103.681082}, {1.349779, 103.681134}, {1.349706, 103.681187}, {1.349633, 103.68124}, {1.349567, 103.681301}, {1.349502, 103.681364}, {1.349438, 103.681427}, {1.349373, 103.681489}, {1.349308, 103.681552}, {1.349251, 103.681622}, {1.349194, 103.681692}, {1.349137, 103.681761}, {1.349078, 103.681829}, {1.349013, 103.681887}, {1.34894, 103.681935}, {1.348921, 103.682019}, {1.348839, 103.682044}, {1.348774, 103.681988}, {1.348776, 103.681902}, {1.34873, 103.681824}, {1.34868, 103.68175}, {1.348635, 103.681672}, {1.348589, 103.681594}, {1.348541, 103.681518}, {1.348495, 103.68144}, {1.34845, 103.681363}, {1.348404, 103.681285}, {1.348358, 103.681207}, {1.348314, 103.681129}, {1.348282, 103.681071}, {1.34825, 103.681014}, {1.34821, 103.680956}, {1.348186, 103.680899}}
	currentPoint := []float64{1.345394, 103.688144}
	currentPointBusStop := []float64{1.35432, 103.68695}
	t.Run("test count distance", func(t *testing.T) {
		res := CountTimeRemaining(currentPoint, currentPointBusStop, point)

		assert.Equal(t, res, 0.0631066719402088)
	})
}
