<p align="center">
<a href="https://dscvit.com">
	<img width="400" src="https://user-images.githubusercontent.com/56252312/159312411-58410727-3933-4224-b43e-4e9b627838a3.png#gh-light-mode-only" alt="GDSC VIT"/>
</a>
	<h2 align="center">Nock Backend</h2>
	<h4 align="center">A team-based attendance app<h4>
</p>

---
[![Join Us](https://img.shields.io/badge/Join%20Us-Developer%20Student%20Clubs-red)](https://dsc.community.dev/vellore-institute-of-technology/)
[![Discord Chat](https://img.shields.io/discord/760928671698649098.svg)](https://discord.gg/498KVdSKWR)

[![DOCS](https://img.shields.io/badge/Documentation-see%20docs-green?style=flat-square&logo=appveyor)](https://documenter.getpostman.com/view/19697822/2s9YR85Yzc) 
  [![UI ](https://img.shields.io/badge/User%20Interface-Link%20to%20UI-orange?style=flat-square&logo=appveyor)](INSERT_UI_LINK_HERE)


## Features
- [x]  Auth via email-password and google/other social login.
- [x]  Create teams and send out invite links.
- [x]  Discover open teams.
- [x]  Join teams and await verification by team admins (if enabled for that team).
- [x]  Get promoted to team admin by team creator/superadmin.
- [x]  Create meetings (attendable by all team members) for a team with locations, timings, etc as a team admin.
- [x]  Get notified of upcoming meetings.
- [x]  Start and end meetings, and take accurate and timebound location-based attendance of members.
- [x]  Other team and meeting admin features.
- [x]  See attendance reports for meetings.
- [ ]  Create meetings for subgroups in teams.

<br>

## Dependencies
 - Go


## Running

### Prerequisites
- Docker

### Directions to Install
```bash
git clone https://github.com/GDGVIT/attendance-app-backend.git
```

### Direction to Execute

```bash
# dev (with live reload)
make dev

# prod
make prod

# remove prod container
make remove-prod
```

## Contributors

<table>
	<tr align="center">
		<td>
		Anirudh Mishra
		<p align="center">
			<img src = "https://avatars.githubusercontent.com/u/91245420" width="150" height="150" alt="Anirudh Mishra">
		</p>
			<p align="center">
				<a href = "https://github.com/anirudhgray">
					<img src = "http://www.iconninja.com/files/241/825/211/round-collaboration-social-github-code-circle-network-icon.svg" width="36" height = "36" alt="GitHub"/>
				</a>
				<a href = "https://www.linkedin.com/in/anirudh-mishra">
					<img src = "http://www.iconninja.com/files/863/607/751/network-linkedin-social-connection-circular-circle-media-icon.svg" width="36" height="36" alt="LinkedIn"/>
				</a>
			</p>
		</td>
	</tr>
</table>

<p align="center">
	Made with ❤ by <a href="https://dscvit.com">DSC VIT</a>
</p>
